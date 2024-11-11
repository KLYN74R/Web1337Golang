package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/KLYN74R/Web1337Golang/pkg/crypto_primitives/tbls"
)

func TestThresholdTransaction(t *testing.T) {
	// Generate random IDs for N parties
	partyCount := uint(5)
	threshold := uint(3)
	partyIDs := tbls.GenerateRandomIds(partyCount)

	// Generate TBLS shares and verification vectors
	verificationVector, sharesForOtherParties := tbls.GenerateTbls(threshold, partyIDs)

	// Verify generated shares for parties
	for i, partyID := range partyIDs {
		isValid := tbls.VerifyShare(partyID, sharesForOtherParties[i], verificationVector)
		if !isValid {
			t.Errorf("Failed to verify share for party %d", i)
		}
	}

	// Derive root public key from verification vectors
	rootPubKey := tbls.DeriveRootPubKey(verificationVector)

	// Create a transaction message
	message := "Sample threshold transaction message"

	// Generate partial signatures for threshold participants
	partialSignatures := []string{}
	idsOfSigners := []string{}
	for i := 0; i < int(threshold); i++ {
		// Use only the share for the current participant
		partialSignature := tbls.GeneratePartialSignature(partyIDs[i], message, []string{sharesForOtherParties[i]})
		partialSignatures = append(partialSignatures, partialSignature)
		idsOfSigners = append(idsOfSigners, partyIDs[i])
	}

	// Build root signature from partial signatures
	rootSignature := tbls.BuildRootSignature(partialSignatures, idsOfSigners)

	// Verify the root signature against the root public key
	isValid := tbls.VerifyRootSignature(rootPubKey, rootSignature, message)
	if !isValid {
		t.Errorf("Failed to verify root signature for threshold transaction")
	}

	// Output Transaction Details
	tx := map[string]interface{}{
		"message":    message,
		"rootPubKey": rootPubKey,
		"signature":  rootSignature,
	}
	jsonData, _ := json.MarshalIndent(tx, "", "  ")
	fmt.Println(string(jsonData))
}
