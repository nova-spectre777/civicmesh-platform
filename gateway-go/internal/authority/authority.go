package authority

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"

	"civicmesh/gateway/internal/domain"
)

type Signer struct {
	KeyID string
	AuthorityID string
	PrivateKey ed25519.PrivateKey
	PublicKey ed25519.PublicKey
}

func Generate(keyID, authorityID string) (*Signer, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil { return nil, err }
	return &Signer{KeyID:keyID,AuthorityID:authorityID,PrivateKey:priv,PublicKey:pub}, nil
}

func NewFromPrivate(keyID, authorityID, privateB64 string) (*Signer, error) {
	raw, err := base64.StdEncoding.DecodeString(privateB64)
	if err != nil { return nil, err }
	if len(raw) != ed25519.PrivateKeySize { return nil, fmt.Errorf("invalid private key size: %d", len(raw)) }
	priv := ed25519.PrivateKey(raw)
	return &Signer{KeyID:keyID,AuthorityID:authorityID,PrivateKey:priv,PublicKey:priv.Public().(ed25519.PublicKey)}, nil
}

func (s *Signer) PublicRecord() domain.AuthorityKey {
	spki, _ := x509.MarshalPKIXPublicKey(s.PublicKey)
	return domain.AuthorityKey{KeyID:s.KeyID,AuthorityID:s.AuthorityID,Algorithm:"Ed25519",PublicKeyB64:base64.StdEncoding.EncodeToString(s.PublicKey),PublicKeySPKIB64:base64.StdEncoding.EncodeToString(spki),Status:"active"}
}

func (s *Signer) SignPayload(envelopeID, payloadJSON string) domain.SignedAlertEnvelope {
	sig := ed25519.Sign(s.PrivateKey, []byte(payloadJSON))
	return domain.SignedAlertEnvelope{EnvelopeID:envelopeID,KeyID:s.KeyID,Algorithm:"Ed25519",PayloadJSON:payloadJSON,SignatureB64:base64.StdEncoding.EncodeToString(sig)}
}

func Verify(env domain.SignedAlertEnvelope, key domain.AuthorityKey) error {
	if key.Status != "active" { return errors.New("authority key is not active") }
	if env.Algorithm != "Ed25519" || key.Algorithm != "Ed25519" { return errors.New("unsupported signature algorithm") }
	if env.KeyID != key.KeyID { return errors.New("key id mismatch") }
	pub, err := base64.StdEncoding.DecodeString(key.PublicKeyB64); if err != nil { return fmt.Errorf("decode public key: %w", err) }
	sig, err := base64.StdEncoding.DecodeString(env.SignatureB64); if err != nil { return fmt.Errorf("decode signature: %w", err) }
	if len(pub) != ed25519.PublicKeySize { return errors.New("invalid public key size") }
	if !ed25519.Verify(ed25519.PublicKey(pub), []byte(env.PayloadJSON), sig) { return errors.New("signature verification failed") }
	return nil
}
