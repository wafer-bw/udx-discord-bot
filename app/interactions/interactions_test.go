package interactions

import (
	"crypto/ed25519"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var interactionsImpl Interactions

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	interactionsImpl = New()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestVerify(t *testing.T) {
	body := "body"
	timestamp := "1500000000"
	pubkey, privkey, err := ed25519.GenerateKey(nil)
	require.Nil(t, err)

	t.Run("success", func(t *testing.T) {
		msg := []byte(timestamp + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize]))

		res := interactionsImpl.Verify([]byte(body), headers, hex.EncodeToString(pubkey))
		require.True(t, res)
	})

	t.Run("failure/modified message parts", func(t *testing.T) {
		msg := []byte(timestamp + "baddata" + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize]))

		res := interactionsImpl.Verify([]byte(body), headers, hex.EncodeToString(pubkey))
		require.False(t, res)
	})

	t.Run("failure/blank signature timestamp", func(t *testing.T) {
		msg := []byte(timestamp + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)

		headers.Set("X-Signature-Timestamp", "")
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize]))

		res := interactionsImpl.Verify([]byte(body), headers, hex.EncodeToString(pubkey))
		require.False(t, res)
	})

	t.Run("failure/blank signature ed25519", func(t *testing.T) {
		headers := http.Header{}

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", "")

		res := interactionsImpl.Verify([]byte(body), headers, hex.EncodeToString(pubkey))
		require.False(t, res)
	})

	t.Run("failure/non-hex public key", func(t *testing.T) {
		msg := []byte(timestamp + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize]))

		res := interactionsImpl.Verify([]byte(body), headers, hex.EncodeToString(pubkey)+"Z")
		require.False(t, res)
	})

	t.Run("failure/non-hex signature", func(t *testing.T) {
		msg := []byte(timestamp + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize])+"Z")

		res := interactionsImpl.Verify([]byte(body), headers, hex.EncodeToString(pubkey))
		require.False(t, res)
	})

	t.Run("failure/wrong length public key", func(t *testing.T) {
		msg := []byte(timestamp + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize]))

		res := interactionsImpl.Verify([]byte(body), headers, hex.EncodeToString(pubkey)+"1111")
		require.False(t, res)
	})

	t.Run("failure/wrong length signature", func(t *testing.T) {
		msg := []byte(timestamp + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize])+"1111")

		res := interactionsImpl.Verify([]byte(body), headers, hex.EncodeToString(pubkey))
		require.False(t, res)
	})
}
