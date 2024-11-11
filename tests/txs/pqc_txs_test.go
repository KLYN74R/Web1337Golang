package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/KLYN74R/Web1337Golang/crypto_primitives/pqc"
)

func TestDilithiumTransaction(t *testing.T) {
	// Generate Dilithium key pair
	pubKey, prvKey := pqc.GenerateDilithiumKeypair()

	var (
		shardID     = "2VEzwUdvSRuv1k2JaAEaMiL7LLNDTUf9bXSapqccCcSb"
		recipient   = "nXSYHp74u88zKPiRi7t22nv4WCBHXUBpGrVw3V93f2s"
		from        = pubKey
		myPrivateKey = prvKey
		message      = fmt.Sprintf("%v%v%v%v%v%v", shardID, from, 0, 0.005, recipient, 0.2)
	)

	// Sign the Transaction
	signature := pqc.GenerateDilithiumSignature(myPrivateKey, message)

	// Verify the Signature
	isValid := pqc.VerifyDilithiumSignature(message, from, signature)
	if !isValid {
		t.Errorf("Failed to verify Dilithium signature for transaction")
	}

	// Output Transaction Details
	tx := map[string]interface{}{
		"from":      from,
		"recipient": recipient,
		"signature": signature,
	}
	jsonData, _ := json.MarshalIndent(tx, "", "  ")
	fmt.Println(string(jsonData))
}

func TestBlissTransaction(t *testing.T) {
	// Generate BLISS key pair
	pubKey, prvKey := pqc.GenerateBlissKeypair()

	var (
		shardID     = "2VEzwUdvSRuv1k2JaAEaMiL7LLNDTUf9bXSapqccCcSb"
		recipient   = "nXSYHp74u88zKPiRi7t22nv4WCBHXUBpGrVw3V93f2s"
		from        = pubKey
		myPrivateKey = prvKey
		message      = fmt.Sprintf("%v%v%v%v%v%v", shardID, from, 0, 0.005, recipient, 0.2)
	)

	// Sign the Transaction
	signature := pqc.GenerateBlissSignature(myPrivateKey, message)

	// Verify the Signature
	isValid := pqc.VerifyBlissSignature(message, from, signature)
	if !isValid {
		t.Errorf("Failed to verify BLISS signature for transaction")
	}

	// Output Transaction Details
	tx := map[string]interface{}{
		"from":      from,
		"recipient": recipient,
		"signature": signature,
	}
	jsonData, _ := json.MarshalIndent(tx, "", "  ")
	fmt.Println(string(jsonData))
}
