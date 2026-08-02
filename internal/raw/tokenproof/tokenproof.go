package tokenproof

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	HeaderName         = "X-Gitpod-CNF-Proof"
	TypeNaClBoxDPoPV1  = "nacl-box-dpop-v1"
	VersionNaClBoxDPoP = "gitpod-cnf-proof-v1"
)

// ConfirmationClaim is the custom RFC 7800-inspired cnf claim carried in
// runner-issued actor access tokens.
type ConfirmationClaim struct {
	KeyID          string `json:"kid,omitempty"`
	ProofType      string `json:"gitpod_proof_type,omitempty"`
	SealedProofKey string `json:"gitpod_sealed_proof_key,omitempty"`
}

// Proof is the request-bound sidecar proof attached by runners when the token
// includes a cnf claim.
type Proof struct {
	Version         string `json:"v"`
	IssuedAt        int64  `json:"iat"`
	ID              string `json:"jti"`
	RPC             string `json:"rpc"`
	AccessTokenHash string `json:"ath"`
	MAC             string `json:"mac"`
}

func AccessTokenHash(token string) string {
	sum := sha256.Sum256([]byte(token))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func CanonicalString(proof Proof) string {
	return fmt.Sprintf("%s\n%s\n%s\n%d\n%s", proof.Version, proof.AccessTokenHash, proof.RPC, proof.IssuedAt, proof.ID)
}

func EncodeProof(proof Proof) (string, error) {
	payload, err := json.Marshal(proof)
	if err != nil {
		return "", fmt.Errorf("marshal proof: %w", err)
	}

	return base64.RawURLEncoding.EncodeToString(payload), nil
}

func DecodeProof(header string) (Proof, error) {
	encoded := strings.TrimSpace(header)
	if encoded == "" {
		return Proof{}, fmt.Errorf("proof header is empty")
	}

	payload, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return Proof{}, fmt.Errorf("decode proof header: %w", err)
	}

	var proof Proof
	if err := json.Unmarshal(payload, &proof); err != nil {
		return Proof{}, fmt.Errorf("unmarshal proof header: %w", err)
	}

	return proof, nil
}
