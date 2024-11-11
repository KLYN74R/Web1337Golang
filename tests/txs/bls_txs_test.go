package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	web1337 "github.com/KLYN74R/Web1337Golang"
	kly_bls "github.com/KLYN74R/Web1337Golang/crypto_primitives/bls"
)

func TestBlsTransaction(t *testing.T) {
	// Generate Keypairs
	secretKey1, publicKey1 := kly_bls.GenerateKeypair()
	secretKey2, publicKey2 := kly_bls.GenerateKeypair()
	_, publicKey3 := kly_bls.GenerateKeypair()

	// Aggregate Public Keys
	rootPubKey := kly_bls.AggregatePubKeys([]string{publicKey1, publicKey2, publicKey3})
	aggregatedPubOfActive := kly_bls.AggregatePubKeys([]string{publicKey1, publicKey2})

	var (
		recipient   = "nXSYHp74u88zKPiRi7t22nv4WCBHXUBpGrVw3V93f2s"
		from        = rootPubKey
		nonce       uint    = 0
		fee         float32 = 0.00005
		amountInKLY float32 = 0.2
		afkSigners  = []string{publicKey3}
	)

	myOptions := web1337.Options{
		SymbioteID:      "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	// Create Multisig Transaction
	defTx := sdkHandler.CreateMultisigTransaction(from, aggregatedPubOfActive, "", afkSigners, nonce, fee, recipient, amountInKLY, nil)

	// Sign the Transaction
	message := fmt.Sprintf("%v%v%v%v%v%v", from, aggregatedPubOfActive, nonce, fee, recipient, amountInKLY)
	signature1 := kly_bls.GenerateSignature(secretKey1, message)
	signature2 := kly_bls.GenerateSignature(secretKey2, message)

	// Aggregate Signatures
	aggregatedSignature := kly_bls.AggregateSignatures([]string{signature1, signature2})
	defTx.Sig = aggregatedSignature

	// Verify the Aggregated Signature
	isValid := kly_bls.VerifySignature(aggregatedPubOfActive, message, aggregatedSignature)
	if !isValid {
		t.Errorf("Failed to verify aggregated signature")
	}

	// Marshal Transaction to JSON and Print
	jsonData, _ := json.MarshalIndent(defTx, "", " ")
	fmt.Println(string(jsonData))
}
