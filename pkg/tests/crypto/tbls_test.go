package tests

import (
	"fmt"
	"testing"

	kly_tbls "github.com/KLYN74R/Web1337Golang/pkg/crypto_primitives/tbls"
)

func TestTBlsProcess(t *testing.T) {

	/*
		T = 2 (Threshold)
		N = 3 (Number of group members)
	*/

	// Generate random IDs for 3 participants
	randomIDs := kly_tbls.GenerateRandomIds(3)
	fmt.Printf("Generated IDs for participants: %v\n\n", randomIDs)

	// Generate verification vectors and secret shares for each participant
	vvec1, secretShares1 := kly_tbls.GenerateTbls(2, randomIDs)
	vvec2, secretShares2 := kly_tbls.GenerateTbls(2, randomIDs)
	vvec3, secretShares3 := kly_tbls.GenerateTbls(2, randomIDs)

	fmt.Printf("Verification Vector 1: %v\n", vvec1)
	fmt.Printf("Verification Vector 2: %v\n", vvec2)
	fmt.Printf("Verification Vector 3: %v\n\n", vvec3)

	// Derive root public key from all verification vectors
	rootPubKey := kly_tbls.DeriveRootPubKey(vvec1, vvec2, vvec3)
	fmt.Printf("Derived Root Public Key: %s\n\n", rootPubKey)

	// Scenario: Members 1 and 2 agree to sign a message, while member 3 disagrees
	msg := "Hello World"

	// Construct secret shares required for generating partial signatures
	secretSharesFor1 := []string{secretShares1[0], secretShares2[0], secretShares3[0]}
	secretSharesFor2 := []string{secretShares1[1], secretShares2[1], secretShares3[1]}

	// Generate partial signatures for members 1 and 2
	partialSignature1 := kly_tbls.GeneratePartialSignature(randomIDs[0], msg, secretSharesFor1)
	partialSignature2 := kly_tbls.GeneratePartialSignature(randomIDs[1], msg, secretSharesFor2)

	fmt.Printf("Partial Signature 1: %s\n", partialSignature1)
	fmt.Printf("Partial Signature 2: %s\n\n", partialSignature2)

	// Aggregate the partial signatures into a root signature
	rootSignature := kly_tbls.BuildRootSignature([]string{partialSignature1, partialSignature2}, []string{randomIDs[0], randomIDs[1]})
	fmt.Printf("Root Signature: %s\n\n", rootSignature)

	// Verify the root signature against the root public key and message
	isSignatureValid := kly_tbls.VerifyRootSignature(rootPubKey, rootSignature, msg)
	if isSignatureValid {
		fmt.Println("Root signature verification successful. The signature is valid.")
	} else {
		t.Error("Root signature verification failed. The signature is invalid.")
	}

	// Negative Test Case: Modify the message and verify the signature (Expected: Invalid)
	alteredMsg := "Hello Altered World"
	isSignatureValidForAlteredMsg := kly_tbls.VerifyRootSignature(rootPubKey, rootSignature, alteredMsg)
	if isSignatureValidForAlteredMsg {
		t.Error("Verification succeeded with an altered message. This should not happen.")
	} else {
		fmt.Println("Verification failed as expected with altered message. The signature is invalid.")
	}

	fmt.Println("\nTest completed.")
}
