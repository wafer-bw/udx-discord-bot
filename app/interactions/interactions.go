package interactions

import (
	"crypto/ed25519"
	"encoding/hex"
	"net/http"
)

// Deps defines `Interactions` dependencies
type Deps struct{}

// impl implements `Interactions` properties
type impl struct{}

// Interactions interfaces `Interactions` methods
type Interactions interface {
	Verify(rawBody []byte, headers http.Header, publicKey string) bool
}

// New returns a new `Interactions` interface
func New() Interactions {
	return &impl{}
}

// Verify verifies that requests from discord are authorized using ed25519
// https://discord.com/developers/docs/interactions/slash-commands#security-and-authorization
func (impl *impl) Verify(rawBody []byte, headers http.Header, publicKey string) bool {
	signature := headers.Get("x-signature-ed25519")
	if signature == "" {
		return false
	}

	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	if len(sig) != ed25519.SignatureSize {
		return false
	}

	timestamp := headers.Get("x-signature-timestamp")
	if timestamp == "" {
		return false
	}

	keyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		return false
	}

	key := ed25519.PublicKey(keyBytes)
	if len(key) != 32 {
		return false
	}

	msg := []byte(timestamp + string(rawBody))
	return ed25519.Verify(key, msg, sig)
}
