package tests

import (
	"fmt"
	"testing"

	kly_bls "github.com/KLYN74R/Web1337Golang/crypto_primitives/bls"
)

func TestBlsProcess(t *testing.T) {

	// Generate keypair

	privateKey, publicKey := kly_bls.GenerateKeypair()

	fmt.Println("Privatekey is => ", privateKey)

	fmt.Println("Publickey is => ", publicKey)

	// Generate signature

	message := "Hello KLY"

	signa := kly_bls.GenerateSignature(privateKey, message)

	fmt.Println("Signa is => ", signa)

	// Now verify (True Positive)
	fmt.Println("Is ok with norm message => ", kly_bls.VerifySignature(publicKey, message, signa))

	// Now verify with wrong msg (True Negative)
	fmt.Println("Is ok with norm message => ", kly_bls.VerifySignature(publicKey, "Hello badass", signa))

	// Now generate more keypairs to test aggregation

	privateKey1, publicKey1 := kly_bls.GenerateKeypair()
	_, publicKey2 := kly_bls.GenerateKeypair()
	_, publicKey3 := kly_bls.GenerateKeypair()

	signa1 := kly_bls.GenerateSignature(privateKey1, message)
	// signa2 := kly_bls.GenerateBlsSignature(privateKey2, message)
	// signa3 := kly_bls.GenerateBlsSignature(privateKey3, message)

	aggregatedSigna := kly_bls.AggregateSignatures([]string{signa, signa1})

	fmt.Println("Aggregated signa is => ", aggregatedSigna)

	// Aggregate pubkeys

	rootPubKey := kly_bls.AggregatePubKeys([]string{publicKey, publicKey1, publicKey2, publicKey3})

	fmt.Println("RootPubKey is => ", rootPubKey)

	// Verify with threshold

	aggregatedPubOfSigners := kly_bls.AggregatePubKeys([]string{publicKey, publicKey1})

	fmt.Println("Aggregated 0 and 1 is => ", aggregatedPubOfSigners)

	aggregatedPub23 := kly_bls.AggregatePubKeys([]string{publicKey2, publicKey3})

	fmt.Println("Aggregated 2 and 3 is => ", aggregatedPub23)

	fmt.Println("Their sum => ", kly_bls.AggregatePubKeys([]string{aggregatedPubOfSigners, aggregatedPub23}))

	fmt.Println("Is threshold reached => ", kly_bls.VerifyThresholdSignature(aggregatedPubOfSigners, aggregatedSigna, rootPubKey, message, []string{publicKey2, publicKey3}, 2))

}
