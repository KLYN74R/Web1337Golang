package tests

import (
	"fmt"
	"testing"

	kly_bls "github.com/KLYN74R/Web1337Golang/pkg/crypto_primitives/bls"
)

func TestBlsProcess(t *testing.T) {
	// Generate keypair for the first signer
	privateKey, publicKey := kly_bls.GenerateKeypair()
	fmt.Printf("Private Key: %s\n", privateKey)
	fmt.Printf("Public Key: %s\n", publicKey)

	// Generate signature for a sample message
	message := "Hello KLY"
	signa := kly_bls.GenerateSignature(privateKey, message)
	fmt.Printf("Signature for message '%s': %s\n", message, signa)

	// Verify signature with correct message (Expected: True)
	if !kly_bls.VerifySignature(publicKey, message, signa) {
		t.Error("Failed to verify signature with correct message.")
	} else {
		fmt.Println("Verification successful with correct message.")
	}

	// Verify signature with incorrect message (Expected: False)
	wrongMessage := "Hello badass"
	if kly_bls.VerifySignature(publicKey, wrongMessage, signa) {
		t.Error("Verification succeeded with incorrect message. This should not happen.")
	} else {
		fmt.Println("Verification failed as expected with incorrect message.")
	}

	// Generate more keypairs to test aggregation
	privateKey1, publicKey1 := kly_bls.GenerateKeypair()
	_, publicKey2 := kly_bls.GenerateKeypair()
	_, publicKey3 := kly_bls.GenerateKeypair()

	fmt.Printf("Public Key 1: %s\n", publicKey1)
	fmt.Printf("Public Key 2: %s\n", publicKey2)
	fmt.Printf("Public Key 3: %s\n", publicKey3)

	// Generate another signature with a different private key
	signa1 := kly_bls.GenerateSignature(privateKey1, message)
	fmt.Printf("Signature 1 for message '%s': %s\n", message, signa1)

	// Aggregate signatures from two signers
	aggregatedSigna := kly_bls.AggregateSignatures([]string{signa, signa1})
	fmt.Printf("Aggregated Signature: %s\n", aggregatedSigna)

	// Aggregate public keys from all four signers
	rootPubKey := kly_bls.AggregatePubKeys([]string{publicKey, publicKey1, publicKey2, publicKey3})
	fmt.Printf("Aggregated Public Key (Root): %s\n", rootPubKey)

	// Aggregate public keys from the first two signers
	aggregatedPubOfSigners := kly_bls.AggregatePubKeys([]string{publicKey, publicKey1})
	fmt.Printf("Aggregated Public Key (Signers 0 and 1): %s\n", aggregatedPubOfSigners)

	// Aggregate public keys from the last two signers
	aggregatedPub23 := kly_bls.AggregatePubKeys([]string{publicKey2, publicKey3})
	fmt.Printf("Aggregated Public Key (Signers 2 and 3): %s\n", aggregatedPub23)

	// Aggregate all public keys to verify they match the root public key
	finalAggregatedPub := kly_bls.AggregatePubKeys([]string{aggregatedPubOfSigners, aggregatedPub23})
	fmt.Printf("Final Aggregated Public Key (All Signers): %s\n", finalAggregatedPub)

	// Verify if threshold signature is valid
	reverseThreshold := uint(2)
	isThresholdReached := kly_bls.VerifyThresholdSignature(
		aggregatedPubOfSigners,
		aggregatedSigna,
		rootPubKey,
		message,
		[]string{publicKey2, publicKey3},
		reverseThreshold,
	)

	if isThresholdReached {
		fmt.Println("Threshold verification succeeded.")
	} else {
		t.Error("Threshold verification failed. Expected successful verification.")
	}
}
