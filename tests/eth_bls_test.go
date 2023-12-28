package tests

import (
	"fmt"
	"testing"

	bls "github.com/herumi/bls-eth-go-binary/bls"
)

func TestEthBls(t *testing.T) {
	// Init lib
	bls.Init(bls.BLS12_381)

	// Generate keys
	secretKey := bls.SecretKey{}
	secretKey.SetByCSPRNG()
	publicKey := secretKey.GetPublicKey()

	message := "Hello, BLS!"

	// Create signature
	signature := secretKey.Sign(message)

	// Вывод результатов
	fmt.Printf("Secret Key: %s\n", secretKey.GetHexString())
	fmt.Printf("Public Key: %s\n", publicKey.SerializeToHexStr())
	fmt.Printf("Message: %s\n", string(message))
	fmt.Printf("Signature: %s\n", signature.SerializeToHexStr())

	// Try to recover pubkey

	bla := bls.PublicKey{}

	fmt.Println("Bla", bla)

	bla.DeserializeHexStr(publicKey.SerializeToHexStr())

	fmt.Println("Now bla is ", bla.SerializeToHexStr())

	// Try to recover signature

	fmt.Println("\n\n===================================")

	blaSig := bls.Sign{}

	fmt.Println("BlaSig", blaSig)

	blaSig.DeserializeHexStr(signature.SerializeToHexStr())

	fmt.Println("Now blaSig is ", blaSig.SerializeToHexStr())

	// Try to recover secret key

	fmt.Println("\n\n===================================")

	blaSecretKey := bls.SecretKey{}

	fmt.Println("BlaSecKey", blaSecretKey)

	blaSecretKey.DeserializeHexStr(secretKey.SerializeToHexStr())

	fmt.Println("Now blaSecKey is ", blaSecretKey.SerializeToHexStr())

	// Check signature
	if signature.Verify(publicKey, message) {
		fmt.Println("Signature is valid.")
	} else {
		fmt.Println("Signature is invalid.")
	}

	// Must be false
	changedMessage := "Modified message"
	if signature.Verify(publicKey, changedMessage) {
		fmt.Println("Modified signature is valid. This should not happen.")
	} else {
		fmt.Println("Modified signature is invalid, as expected.")
	}

	// Now try to aggregate public keys & signatures

	secretKey2 := bls.SecretKey{}

	secretKey2.SetByCSPRNG()

	publicKey2 := secretKey2.GetPublicKey()

	// Create signature
	signature2 := secretKey2.Sign(message)

	// Get the rootpub from pubkeys 1 and 2

	rootPubKey := bls.PublicKey{}

	rootPubKey.Add(publicKey)
	rootPubKey.Add(publicKey2)

	fmt.Println("RootPubKey is => ", rootPubKey.SerializeToHexStr())

	// Get the aggregated signature

	aggregatedSignature := bls.Sign{}

	// aggregatedSignature.Add(signature)
	// aggregatedSignature.Add(signature2)

	aggregatedSignature.Aggregate([]bls.Sign{*signature, *signature2})

	fmt.Println("Aggregated signa is => ", aggregatedSignature.SerializeToHexStr())

	// Now verify the aggregated signature with rootpub

	fmt.Println("Is aggregated signature ok ? =>", aggregatedSignature.Verify(&rootPubKey, message))

}
