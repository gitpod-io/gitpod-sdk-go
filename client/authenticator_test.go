package client

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	jose "github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/google/go-cmp/cmp"
)

func TestIssuerFromApiEndpoint(t *testing.T) {
	type Expectation struct {
		Error  string
		Result string
	}
	tests := []struct {
		Input       string
		Expectation Expectation
	}{
		{Input: "https://example.com/api", Expectation: Expectation{Result: "https://example.com"}},
		{Input: "https://example.com/some/place/with/a/path/api", Expectation: Expectation{Result: "https://example.com"}},
		{Input: "not-a-valid-url", Expectation: Expectation{Error: "invalid API endpoint: not-a-valid-url"}},
	}

	for _, test := range tests {
		t.Run(test.Input, func(t *testing.T) {
			var act Expectation

			res, err := IssuerFromApiEndpoint(context.Background(), test.Input)
			if err != nil {
				act.Error = err.Error()
			}
			act.Result = res

			if diff := cmp.Diff(test.Expectation, act); diff != "" {
				t.Errorf("IssuerFromApiEndpoint() mismatch (-want  got):\n%s", diff)
			}
		})
	}
}

func TestAuthenticate(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Errorf("rsa.GenerateKey() error = %v", err)
	}
	const keyID = "k0"
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: key}, &jose.SignerOptions{
		ExtraHeaders: map[jose.HeaderKey]interface{}{
			jose.HeaderKey("kid"): keyID,
		},
	})
	if err != nil {
		t.Errorf("jose.NewSigner() error = %v", err)
	}
	signToken := func(claims any) string {
		res, err := jwt.Signed(signer).Claims(claims).Serialize()
		if err != nil {
			t.Fatal(err)
		}
		return res
	}
	const issuer = "foo"
	jwks := jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{
				Key:       key.Public(),
				KeyID:     keyID,
				Algorithm: string(jose.RS256),
				Use:       "sig",
			},
		},
	}

	type Expectation struct {
		Error  string
		Result *AuthenticationResult
	}
	tests := []struct {
		Name        string
		Expectation Expectation
		JWKS        jose.JSONWebKeySet
		Token       string
	}{
		{
			Name: "valid token",
			JWKS: jwks,
			Token: signToken(jwt.Claims{
				Issuer:  issuer,
				Subject: "user/bla",
			}),
			Expectation: Expectation{
				Result: &AuthenticationResult{
					Claims:      jwt.Claims{Issuer: "foo", Subject: "user/bla"},
					Principal:   "user",
					PrincipalID: "bla",
				},
			},
		},
		{
			Name: "unknown issuer",
			Token: signToken(jwt.Claims{
				Issuer:  "bar",
				Subject: "user/bla",
			}),
			Expectation: Expectation{
				Error: "unknown issuer: bar",
			},
		},
		{
			Name: "unknown key",
			JWKS: jose.JSONWebKeySet{
				Keys: []jose.JSONWebKey{
					{
						Key:       key.Public(),
						KeyID:     "k1",
						Algorithm: string(jose.RS256),
						Use:       "sig",
					},
				},
			},
			Token: signToken(jwt.Claims{
				Issuer:  issuer,
				Subject: "user/bla",
			}),
			Expectation: Expectation{
				Error: "unknown key: k0",
			},
		},
		{
			Name: "no key header",
			JWKS: jwks,
			Token: func() string {
				signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: key}, nil)
				if err != nil {
					t.Errorf("jose.NewSigner() error = %v", err)
				}
				res, err := jwt.Signed(signer).Claims(jwt.Claims{
					Issuer:  issuer,
					Subject: "user/bla",
				}).Serialize()
				if err != nil {
					t.Fatal(err)
				}
				return res
			}(),
			Expectation: Expectation{
				Error: "no key header",
			},
		},
		{
			Name: "invalid subject",
			JWKS: jwks,
			Token: signToken(jwt.Claims{
				Issuer:  issuer,
				Subject: "user",
			}),
			Expectation: Expectation{
				Error: "invalid subject: user",
			},
		},
		{
			Name:  "invalid token",
			Token: "foo",
			Expectation: Expectation{
				Error: "cannot parse token: go-jose/go-jose: compact JWS format must have three parts",
			},
		},
		{
			Name: "signature mismatch",
			JWKS: jwks,
			Token: func() string {
				key, err := rsa.GenerateKey(rand.Reader, 2048)
				if err != nil {
					t.Errorf("rsa.GenerateKey() error = %v", err)
				}
				signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: key}, &jose.SignerOptions{
					ExtraHeaders: map[jose.HeaderKey]interface{}{
						jose.HeaderKey("kid"): keyID,
					},
				})
				if err != nil {
					t.Errorf("jose.NewSigner() error = %v", err)
				}
				res, err := jwt.Signed(signer).Claims(jwt.Claims{
					Issuer:  issuer,
					Subject: "user/bla",
				}).Serialize()
				if err != nil {
					t.Fatal(err)
				}
				return res
			}(),
			Expectation: Expectation{
				Error: "cannot validate token: go-jose/go-jose: error in cryptographic primitive",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var act Expectation

			authenticator := OIDCAuthenticator{
				issuers: map[string]jose.JSONWebKeySet{
					"foo": test.JWKS,
				},
			}
			res, err := authenticator.Authenticate(context.Background(), test.Token)
			if err != nil {
				act.Error = err.Error()
			}
			act.Result = res

			if diff := cmp.Diff(test.Expectation, act); diff != "" {
				t.Errorf("Authenticate() mismatch (-want  got):\n%s", diff)
			}
		})
	}
}

func TestRefreshIssuerJWKS(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("rsa.GenerateKey() error = %v", err)
	}

	type Expectation struct {
		Error  string
		Issuer map[string]jose.JSONWebKeySet
	}
	tests := []struct {
		Name             string
		Expectation      Expectation
		JWKS             jose.JSONWebKeySet
		RawJWKS          string
		DiscoveryHandler func(baseURL string) http.Handler
	}{
		{
			Name: "valid JWKS",
			JWKS: jose.JSONWebKeySet{
				Keys: []jose.JSONWebKey{
					{
						Key:       key.Public(),
						KeyID:     "k0",
						Algorithm: string(jose.RS256),
						Use:       "sig",
					},
				},
			},
			Expectation: Expectation{
				Issuer: map[string]jose.JSONWebKeySet{
					"example.com": {
						Keys: []jose.JSONWebKey{
							{
								Key:                         key.Public(),
								KeyID:                       "k0",
								Algorithm:                   string(jose.RS256),
								Use:                         "sig",
								Certificates:                []*x509.Certificate{},
								CertificateThumbprintSHA1:   []uint8{},
								CertificateThumbprintSHA256: []uint8{},
							},
						},
					},
				},
			},
		},
		{
			Name:    "invalid JWKS",
			RawJWKS: `{"keys":[{"kty":"foobar","alg":"unknown","use":"sig","kid":"k0","n":"foo","e":"AQAB"}]}`,
			Expectation: Expectation{
				Error:  "cannot decode JWKS from example.com/jwks: go-jose/go-jose: unknown json web key type 'foobar'",
				Issuer: map[string]jose.JSONWebKeySet{},
			},
		},
		{
			Name: "JWKS not found",
			DiscoveryHandler: func(baseURL string) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					_, _ = w.Write([]byte(`{"issuer":"` + baseURL + `","jwks_uri":"` + baseURL + "/not-found" + `"}`))
				})
			},
			Expectation: Expectation{
				Error:  "cannot decode JWKS from example.com/not-found: json: cannot unmarshal number into Go value of type jose.JSONWebKeySet",
				Issuer: map[string]jose.JSONWebKeySet{},
			},
		},
		{
			Name: "discovery error",
			DiscoveryHandler: func(baseURL string) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				})
			},
			Expectation: Expectation{
				Error: `cannot discover OIDC configuration for example.com: OpenID Provider Configuration Discovery has failed
http status not ok: 500 Internal Server Error `,
				Issuer: map[string]jose.JSONWebKeySet{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var act Expectation

			mux := http.NewServeMux()
			srv := httptest.NewServer(mux)
			t.Cleanup(srv.Close)
			if test.DiscoveryHandler != nil {
				mux.Handle("/.well-known/openid-configuration", test.DiscoveryHandler(srv.URL))
			} else {
				mux.Handle("/.well-known/openid-configuration", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					_, _ = w.Write([]byte(`{"issuer":"` + srv.URL + `","jwks_uri":"` + srv.URL + `/jwks"}`))
				}))
			}
			if test.RawJWKS != "" {
				mux.Handle("/jwks", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					_, _ = w.Write([]byte(test.RawJWKS))
				}))
			} else {
				mux.Handle("/jwks", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					err := json.NewEncoder(w).Encode(test.JWKS)
					if err != nil {
						t.Fatal(err)
					}
				}))
			}

			authenticator := OIDCAuthenticator{
				issuers: make(map[string]jose.JSONWebKeySet),
			}
			err := authenticator.refreshIssuerJWKS(context.Background(), srv.URL)
			if err != nil {
				act.Error = strings.ReplaceAll(err.Error(), srv.URL, "example.com")
			}
			act.Issuer = make(map[string]jose.JSONWebKeySet, len(authenticator.issuers))
			for k, v := range authenticator.issuers {
				act.Issuer[strings.ReplaceAll(k, srv.URL, "example.com")] = v
			}

			if diff := cmp.Diff(test.Expectation, act); diff != "" {
				t.Errorf("refreshIssuerJWKS() mismatch (-want  got):\n%s", diff)
			}
		})
	}
}
