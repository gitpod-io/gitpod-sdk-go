package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	jose "github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/zitadel/oidc/v3/pkg/client"
)

// IssuerFromApiEndpoint extracts the issuer URL from a control plane API endpoint.
func IssuerFromApiEndpoint(ctx context.Context, apiEndpoint string) (string, error) {
	// API endpoint is expected to be in the form of https://<issuer>/api.
	// Hence, this function expects that the issuer does not have path component.
	//
	// This function takes a context and returns an error so that we can, instead of spreading the convention,
	// we can query the API instead.

	apiURL, err := url.Parse(apiEndpoint)
	if err != nil {
		return "", fmt.Errorf("cannot parse API endpoint: %w", err)
	}
	if apiURL.Scheme == "" || apiURL.Host == "" {
		return "", fmt.Errorf("invalid API endpoint: %s", apiEndpoint)
	}
	return apiURL.Scheme + "://" + apiURL.Host, nil
}

// NewOIDCAuthenticator creates a new OIDC authenticator
func NewOIDCAuthenticator(validIssuers []string) (*OIDCAuthenticator, error) {
	res := &OIDCAuthenticator{
		issuers: make(map[string]jose.JSONWebKeySet),
	}

	for _, issuer := range validIssuers {
		err := res.refreshIssuerJWKS(context.Background(), issuer)
		if err != nil {
			return nil, fmt.Errorf("cannot get JWKS for %s: %w", issuer, err)
		}
	}

	return res, nil
}

// OIDCAuthenticator is an authenticator that validates OIDC tokens
type OIDCAuthenticator struct {
	issuers map[string]jose.JSONWebKeySet
	mu      sync.RWMutex
}

func (a *OIDCAuthenticator) refreshIssuerJWKS(ctx context.Context, issuer string) error {
	timeout := 30 * time.Second
	if dl, ok := ctx.Deadline(); ok {
		timeout = time.Until(dl)
	}

	httpClient := &http.Client{
		Timeout: timeout,
	}

	discovery, err := client.Discover(ctx, issuer, httpClient)
	if err != nil {
		return fmt.Errorf("cannot discover OIDC configuration for %s: %w", issuer, err)
	}
	if discovery.Issuer != issuer {
		return fmt.Errorf("discovered issuer %s does not match expected issuer %s", discovery.Issuer, issuer)
	}

	jwksResp, err := httpClient.Get(discovery.JwksURI)
	if err != nil {
		return fmt.Errorf("cannot fetch JWKS from %s: %w", discovery.JwksURI, err)
	}
	defer jwksResp.Body.Close()

	var jwks jose.JSONWebKeySet
	err = json.NewDecoder(jwksResp.Body).Decode(&jwks)
	if err != nil {
		return fmt.Errorf("cannot decode JWKS from %s: %w", discovery.JwksURI, err)
	}

	a.mu.Lock()
	a.issuers[issuer] = jwks
	a.mu.Unlock()
	return nil
}

func (a *OIDCAuthenticator) Authenticate(ctx context.Context, token string) (*AuthenticationResult, error) {
	tkn, err := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.RS256, jose.RS384, jose.RS512, jose.ES256, jose.ES384, jose.ES512})
	if err != nil {
		return nil, fmt.Errorf("cannot parse token: %w", err)
	}
	if len(tkn.Headers) == 0 {
		return nil, fmt.Errorf("no headers in token")
	}
	hdr := tkn.Headers[0]
	if hdr.KeyID == "" {
		return nil, fmt.Errorf("no key header")
	}

	var unsafeClaims jwt.Claims
	err = tkn.UnsafeClaimsWithoutVerification(&unsafeClaims)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal claims: %w", err)
	}
	a.mu.RLock()
	keyset, ok := a.issuers[unsafeClaims.Issuer]
	a.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unknown issuer: %s", unsafeClaims.Issuer)
	}
	keys := keyset.Key(hdr.KeyID)
	if len(keys) == 0 {
		return nil, fmt.Errorf("unknown key: %s", hdr.KeyID)
	}

	var safeClaims struct {
		jwt.Claims
		Organization string `json:"org"`
	}
	err = tkn.Claims(keys[0].Key, &safeClaims)
	if err != nil {
		return nil, fmt.Errorf("cannot validate token: %w", err)
	}

	segs := strings.Split(safeClaims.Subject, "/")
	if len(segs) != 2 {
		return nil, fmt.Errorf("invalid subject: %s", safeClaims.Subject)
	}

	result := &AuthenticationResult{
		Claims:         safeClaims.Claims,
		OrganizationID: safeClaims.Organization,
		Principal:      segs[0],
		PrincipalID:    segs[1],
	}
	return result, nil
}

// AuthenticationResult contains the result of an authentication process
type AuthenticationResult struct {
	Claims         jwt.Claims
	OrganizationID string
	Principal      string
	PrincipalID    string
}
