package auth

import (
	"crypto/ed25519"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/mocks"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
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
		conf := mocks.GetConf()
		conf.Credentials.PublicKey = hex.EncodeToString(pubkey)
		authorizationImpl := New(&Deps{}, conf)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize]))

		res := authorizationImpl.Verify([]byte(body), headers)
		require.True(t, res)
	})

	t.Run("failure/modified message parts", func(t *testing.T) {
		msg := []byte(timestamp + "baddata" + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)
		conf := mocks.GetConf()
		conf.Credentials.PublicKey = hex.EncodeToString(pubkey)
		authorizationImpl := New(&Deps{}, conf)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize]))

		res := authorizationImpl.Verify([]byte(body), headers)
		require.False(t, res)
	})

	t.Run("failure/blank signature timestamp", func(t *testing.T) {
		msg := []byte(timestamp + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)
		conf := mocks.GetConf()
		conf.Credentials.PublicKey = hex.EncodeToString(pubkey)
		authorizationImpl := New(&Deps{}, conf)

		headers.Set("X-Signature-Timestamp", "")
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize]))

		res := authorizationImpl.Verify([]byte(body), headers)
		require.False(t, res)
	})

	t.Run("failure/blank signature ed25519", func(t *testing.T) {
		headers := http.Header{}
		conf := mocks.GetConf()
		conf.Credentials.PublicKey = hex.EncodeToString(pubkey)
		authorizationImpl := New(&Deps{}, conf)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", "")

		res := authorizationImpl.Verify([]byte(body), headers)
		require.False(t, res)
	})

	t.Run("failure/non-hex public key", func(t *testing.T) {
		msg := []byte(timestamp + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)
		conf := mocks.GetConf()
		conf.Credentials.PublicKey = hex.EncodeToString(pubkey) + "Z"
		authorizationImpl := New(&Deps{}, conf)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize]))

		res := authorizationImpl.Verify([]byte(body), headers)
		require.False(t, res)
	})

	t.Run("failure/non-hex signature", func(t *testing.T) {
		msg := []byte(timestamp + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)
		conf := mocks.GetConf()
		conf.Credentials.PublicKey = hex.EncodeToString(pubkey)
		authorizationImpl := New(&Deps{}, conf)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize])+"Z")

		res := authorizationImpl.Verify([]byte(body), headers)
		require.False(t, res)
	})

	t.Run("failure/wrong length public key", func(t *testing.T) {
		msg := []byte(timestamp + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)
		conf := mocks.GetConf()
		conf.Credentials.PublicKey = hex.EncodeToString(pubkey) + "1111"
		authorizationImpl := New(&Deps{}, conf)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize]))

		res := authorizationImpl.Verify([]byte(body), headers)
		require.False(t, res)
	})

	t.Run("failure/wrong length signature", func(t *testing.T) {
		msg := []byte(timestamp + body)
		headers := http.Header{}
		signature := ed25519.Sign(privkey, msg)
		conf := mocks.GetConf()
		conf.Credentials.PublicKey = hex.EncodeToString(pubkey)
		authorizationImpl := New(&Deps{}, conf)

		headers.Set("X-Signature-Timestamp", timestamp)
		headers.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize])+"1111")

		res := authorizationImpl.Verify([]byte(body), headers)
		require.False(t, res)
	})
}
