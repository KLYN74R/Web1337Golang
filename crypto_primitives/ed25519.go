package crypto_primitives

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
)

func GenerateEd25519KeyPair() (string, string) {

	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)

	return base58.Encode(publicKey), hex.EncodeToString(privateKey[:32])

}

// Returns signature in base64(to use it in transaction later)

func GenerateEd25519Signature(privateKey, msg string) string {

	privateKeyAsBytes, _ := hex.DecodeString(privateKey)

	privateKeyFromSeed := ed25519.NewKeyFromSeed(privateKeyAsBytes)

	msgAsBytes := []byte(msg)

	return base64.StdEncoding.EncodeToString(ed25519.Sign(privateKeyFromSeed, msgAsBytes))

}

/*
0 - message that was signed
1 - pubKey
2 - signature
*/
func VerifyEd25519Signature(stringMessage, base58PubKey, base64Signature string) bool {

	// Decode evrything

	msgAsBytes := []byte(stringMessage)

	publicKey := base58.Decode(base58PubKey)

	signature, _ := base64.StdEncoding.DecodeString(base64Signature)

	return ed25519.Verify(publicKey, msgAsBytes, signature)

}
