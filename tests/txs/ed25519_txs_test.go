package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	web1337 "github.com/KLYN74R/Web1337Golang"
	ed25519 "github.com/KLYN74R/Web1337Golang/crypto_primitives/ed25519"
)

func TestEd25519ToEd25519Transaction(t *testing.T) {
	myKeypair := ed25519.Ed25519Box{
		Mnemonic:  "smoke suggest security index situate almost ethics tone wash crystal debris mosquito pony extra husband elder over relax width occur inspire keen sudden average",
		Bip44Path: []uint32{44, 7331, 0, 0},
		Pub:       "8LZ1XWiLHwuEVsWfaWmY2sm98hYrFAbGqfi4zwbApfJT",
		Prv:       "MC4CAQAwBQYDK2VwBCIEIHmKXvMPju2bdhbhqUYR08ujs/SYH/WPC3cPQfhnym89",
	}

	var (
		shardID     = "2VEzwUdvSRuv1k2JaAEaMiL7LLNDTUf9bXSapqccCcSb"
		recipient   = "nXSYHp74u88zKPiRi7t22nv4WCBHXUBpGrVw3V93f2s"
		from        = myKeypair.Pub
		myPrivateKey = myKeypair.Prv
		nonce       = uint(0)
		fee         float32 = 0.005
		amountInKLY float32 = 0.2
	)

	myOptions := web1337.Options{
		SymbioteID:      "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	// Create Default Transaction
	defTx, err := sdkHandler.CreateDefaultTransaction(shardID, from, myPrivateKey, nonce, fee, recipient, amountInKLY, nil)
	if err != nil {
		t.Errorf("Failed to create default transaction: %v", err)
	}

	// Sign the Transaction
	message := fmt.Sprintf("%v%v%v%v%v%v", shardID, from, nonce, fee, recipient, amountInKLY)
	signature := ed25519.GenerateSignature(myPrivateKey, message)
	defTx.Sig = signature

	// Verify the Signature
	isValid := ed25519.VerifySignature(message, from, signature)
	if !isValid {
		t.Errorf("Failed to verify signature for default transaction")
	}

	// Marshal Transaction to JSON and Print
	jsonData, _ := json.MarshalIndent(defTx, "", "  ")
	fmt.Println(string(jsonData))

	// Now create Ed25519 to Multisig Transaction
	msigRecipient := "7GPupbq1vtKUgaqVeHiDbEJcxS7sSjwPnbht4eRaDBAEJv8ZKHNCSu2Am3CuWnHjta"
	var rev_t int = 0

	msigTx, err := sdkHandler.CreateDefaultTransaction(shardID, from, myPrivateKey, nonce, fee, msigRecipient, amountInKLY, &rev_t)
	if err != nil {
		t.Errorf("Failed to create multisig transaction: %v", err)
	}

	// Sign the Multisig Transaction
	msigMessage := fmt.Sprintf("%v%v%v%v%v%v%v", shardID, from, nonce, fee, msigRecipient, amountInKLY, rev_t)
	signatureMsig := ed25519.GenerateSignature(myPrivateKey, msigMessage)
	msigTx.Sig = signatureMsig

	// Verify the Multisig Signature
	isValidMsig := ed25519.VerifySignature(msigMessage, from, signatureMsig)
	if !isValidMsig {
		t.Errorf("Failed to verify signature for multisig transaction")
	}

	// Marshal Multisig Transaction to JSON and Print
	jsonDataMsig, _ := json.MarshalIndent(msigTx, "", "  ")
	fmt.Println(string(jsonDataMsig))
}
