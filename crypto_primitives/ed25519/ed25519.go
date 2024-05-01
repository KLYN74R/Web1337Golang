package ed25519

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type Ed25519Box struct {
	Mnemonic  string
	Bip44Path []uint32
	Pub, Prv  string
}

func GenerateKeyPair(mnemonic, mnemonicPassword string, bip44DerivePath []uint32) Ed25519Box {

	if mnemonic == "" {

		// Generate mnemonic if no pre-set

		entropy, _ := bip39.NewEntropy(256)

		mnemonic, _ = bip39.NewMnemonic(entropy)

	}

	// Now generate seed from 24-word mnemonic phrase (24 words = 256 bit security)
	// Seed has 64 bytes
	seed := bip39.NewSeed(mnemonic, mnemonicPassword) // password might be ""(empty) but it's not recommended

	// Generate master keypair from seed

	masterPrivateKey, _ := bip32.NewMasterKey(seed)

	// Now, to derive appropriate keypair - run the cycle over uint32 path-milestones and derive child keypairs

	// In case bip44Path empty - set the default one

	if len(bip44DerivePath) == 0 {

		bip44DerivePath = []uint32{44, 7331, 0, 0}

	}

	// Start derivation from master private key
	var childKey *bip32.Key = masterPrivateKey

	for pathPart := range bip44DerivePath {

		childKey, _ = childKey.NewChildKey(bip32.FirstHardenedChild + uint32(pathPart))

	}

	// Now, based on this - get the appropriate keypair

	publicKeyAsBytes, privateKeyAsBytes := generateKeyPairFromSeed(childKey.Key)

	return Ed25519Box{Mnemonic: mnemonic, Bip44Path: bip44DerivePath, Pub: base58.Encode(publicKeyAsBytes), Prv: hex.EncodeToString(privateKeyAsBytes)}

}

// Returns signature in base64(to use it in transaction later)

func GenerateSignature(privateKey, msg string) string {

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
func VerifySignature(stringMessage, base58PubKey, base64Signature string) bool {

	// Decode evrything

	msgAsBytes := []byte(stringMessage)

	publicKey := base58.Decode(base58PubKey)

	signature, _ := base64.StdEncoding.DecodeString(base64Signature)

	return ed25519.Verify(publicKey, msgAsBytes, signature)

}

// Private inner function

func generateKeyPairFromSeed(seed []byte) ([]byte, []byte) {

	publicKey := make([]byte, 32)

	privateKey := ed25519.NewKeyFromSeed(seed)

	copy(publicKey, privateKey[32:])

	return publicKey, privateKey[:32]

}
