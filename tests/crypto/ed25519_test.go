package tests

import (
	"fmt"
	"testing"

	ed25519 "github.com/KLYN74R/Web1337Golang/crypto_primitives/ed25519"
)

func TestEd25519(t *testing.T) {
	// Define mnemonic and derivation path for key generation
	mnemonic := "smoke suggest security index situate almost ethics tone wash crystal debris mosquito pony extra husband elder over relax width occur inspire keen sudden average"
	mnemonicPassword := ""
	bip44Path := []uint32{44, 7331, 0, 0}

	// Generate keypair using mnemonic and derivation path
	keypair := ed25519.GenerateKeyPair(mnemonic, mnemonicPassword, bip44Path)

	fmt.Printf("Generated Keypair:\n  Public Key: %s\n  Private Key: %s\n", keypair.Pub, keypair.Prv)

	// Create a message to be signed
	message := "BlaBlaBla"

	// Generate signature for the message
	signature := ed25519.GenerateSignature(keypair.Prv, message)
	fmt.Printf("Generated Signature for message '%s': %s\n", message, signature)

	// Verify the signature with the correct message (Expected: True)
	if !ed25519.VerifySignature(message, keypair.Pub, signature) {
		t.Error("Failed to verify signature with correct message.")
	} else {
		fmt.Println("Verification successful with correct message.")
	}

	// Verify the signature with an incorrect message (Expected: False)
	wrongMessage := "WrongMessage"
	if ed25519.VerifySignature(wrongMessage, keypair.Pub, signature) {
		t.Error("Verification succeeded with incorrect message. This should not happen.")
	} else {
		fmt.Println("Verification failed as expected with incorrect message.")
	}
}
