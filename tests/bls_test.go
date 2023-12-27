package tests

import (
	"fmt"
	"testing"

	crypto_primitives "github.com/KLYN74R/Web1337Golang/crypto_primitives"
)

func TestBlsProcess(t *testing.T) {

	// Generate keypair

	privateKey, publicKey := crypto_primitives.GenerateBlsKeypair()

	fmt.Println("Privatekey is => ", privateKey)

	fmt.Println("Publickey is => ", publicKey)

	// Generate signature

	message := "Hello KLY"

	signa := crypto_primitives.GenerateBlsSignature(privateKey, message)

	fmt.Println("Signa is => ", signa)

	// Now verify (True Positive)
	fmt.Println("Is ok with norm message => ", crypto_primitives.VerifyBlsSignature(publicKey, message, signa))

	// Now verify with wrong msg (True Negative)
	fmt.Println("Is ok with norm message => ", crypto_primitives.VerifyBlsSignature(publicKey, "Hello badass", signa))

	// Now generate more keypairs to test aggregation

	privateKey1, publicKey1 := crypto_primitives.GenerateBlsKeypair()
	_, publicKey2 := crypto_primitives.GenerateBlsKeypair()
	_, publicKey3 := crypto_primitives.GenerateBlsKeypair()

	signa1 := crypto_primitives.GenerateBlsSignature(privateKey1, message)
	// signa2 := crypto_primitives.GenerateBlsSignature(privateKey2, message)
	// signa3 := crypto_primitives.GenerateBlsSignature(privateKey3, message)

	aggregatedSigna := crypto_primitives.AggregateBlsSignatures([]string{signa, signa1})

	fmt.Println("Aggregated signa is => ", aggregatedSigna)

	// Aggregate pubkeys

	rootPubKey := crypto_primitives.AggregateBlsPubKeys([]string{publicKey, publicKey1, publicKey2, publicKey3})

	fmt.Println("RootPubKey is => ", rootPubKey)

	// Verify with threshold

	aggregatedPubOfSigners := crypto_primitives.AggregateBlsPubKeys([]string{publicKey, publicKey1})

	fmt.Println("Aggregated 0 and 1 is => ", aggregatedPubOfSigners)

	aggregatedPub23 := crypto_primitives.AggregateBlsPubKeys([]string{publicKey2, publicKey3})

	fmt.Println("Aggregated 2 and 3 is => ", aggregatedPub23)

	fmt.Println("Their sum => ", crypto_primitives.AggregateBlsPubKeys([]string{aggregatedPubOfSigners, aggregatedPub23}))

	fmt.Println("Is threshold reached => ", crypto_primitives.VerifyBlsThresholdSignature(aggregatedPubOfSigners, aggregatedSigna, rootPubKey, message, []string{publicKey2, publicKey3}, 2))

}
