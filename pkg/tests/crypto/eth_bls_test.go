package tests

import (
	"fmt"
	"testing"

	bls "github.com/herumi/bls-eth-go-binary/bls"
)

func TestEthBls(t *testing.T) {
	// Initialize the BLS library for the BLS12-381 curve
	bls.Init(bls.BLS12_381)

	// Generate key pair
	secretKey := bls.SecretKey{}
	secretKey.SetByCSPRNG()
	publicKey := secretKey.GetPublicKey()

	message := "Hello, BLS!"

	// Create signature
	signature := secretKey.Sign(message)

	// Print generated keys and signature
	fmt.Printf("Secret Key: %s\n", secretKey.GetHexString())
	fmt.Printf("Public Key: %s\n", publicKey.SerializeToHexStr())
	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Signature: %s\n", signature.SerializeToHexStr())

	// Recover and verify public key from serialized value
	recoveredPubKey := bls.PublicKey{}
	err := recoveredPubKey.DeserializeHexStr(publicKey.SerializeToHexStr())
	if err != nil {
		t.Errorf("Failed to deserialize public key: %v", err)
	}
	fmt.Printf("Recovered Public Key: %s\n", recoveredPubKey.SerializeToHexStr())

	// Recover and verify signature from serialized value
	recoveredSignature := bls.Sign{}
	err = recoveredSignature.DeserializeHexStr(signature.SerializeToHexStr())
	if err != nil {
		t.Errorf("Failed to deserialize signature: %v", err)
	}
	fmt.Printf("Recovered Signature: %s\n", recoveredSignature.SerializeToHexStr())

	// Recover and verify secret key from serialized value
	recoveredSecretKey := bls.SecretKey{}
	err = recoveredSecretKey.DeserializeHexStr(secretKey.SerializeToHexStr())
	if err != nil {
		t.Errorf("Failed to deserialize secret key: %v", err)
	}
	fmt.Printf("Recovered Secret Key: %s\n", recoveredSecretKey.GetHexString())

	// Verify the original signature
	if signature.Verify(publicKey, message) {
		fmt.Println("Signature is valid.")
	} else {
		t.Error("Signature verification failed.")
	}

	// Verify the signature with a modified message (expected to fail)
	changedMessage := "Modified message"
	if signature.Verify(publicKey, changedMessage) {
		t.Error("Modified signature is valid. This should not happen.")
	} else {
		fmt.Println("Modified signature is invalid, as expected.")
	}

	// Generate a second key pair for aggregation testing
	secretKey2 := bls.SecretKey{}
	secretKey2.SetByCSPRNG()
	publicKey2 := secretKey2.GetPublicKey()
	signature2 := secretKey2.Sign(message)

	fmt.Printf("Second Secret Key: %s\n", secretKey2.GetHexString())
	fmt.Printf("Second Public Key: %s\n", publicKey2.SerializeToHexStr())

	// Aggregate public keys
	aggregatedPubKey := bls.PublicKey{}
	aggregatedPubKey.Add(publicKey)
	aggregatedPubKey.Add(publicKey2)
	fmt.Printf("Aggregated Public Key: %s\n", aggregatedPubKey.SerializeToHexStr())

	// Aggregate signatures
	aggregatedSignature := bls.Sign{}
	aggregatedSignature.Aggregate([]bls.Sign{*signature, *signature2})
	fmt.Printf("Aggregated Signature: %s\n", aggregatedSignature.SerializeToHexStr())

	// Verify the aggregated signature with the aggregated public key
	if aggregatedSignature.Verify(&aggregatedPubKey, message) {
		fmt.Println("Aggregated signature is valid.")
	} else {
		t.Error("Aggregated signature verification failed.")
	}
}
