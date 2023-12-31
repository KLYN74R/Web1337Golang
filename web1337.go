/*

Web1337 by KLY

For Golang devs

*/

package web1337

import (
	"fmt"

	bls "github.com/KLYN74R/Web1337Golang/crypto_primitives/bls"
	ed25519 "github.com/KLYN74R/Web1337Golang/crypto_primitives/ed25519"
	pqc "github.com/KLYN74R/Web1337Golang/crypto_primitives/pqc"
)

func Ed25519Process() bool {

	myPubKey, myPrivateKey := ed25519.GenerateKeyPair()

	fmt.Println("PubKey is ", myPubKey)

	fmt.Println("PrivateKey is ", myPrivateKey)

	signa := ed25519.GenerateSignature(myPrivateKey, "Hello KLY")

	fmt.Println("Signa is ", signa)

	isOk := ed25519.VerifySignature("Hello KLY", myPubKey, signa)

	fmt.Println("Is ok =>", isOk)

	return isOk

}

func BlissProcess() bool {

	myPubKey, myPrivateKey := pqc.GenerateBlissKeypair()

	fmt.Println("PubKey is ", myPubKey)

	fmt.Println("PrivateKey is ", myPrivateKey)

	signa := pqc.GenerateBlissSignature(myPrivateKey, "Hello KLY")

	fmt.Println("Signa is ", signa)

	isOk := pqc.VerifyBlissSignature("Hello KLY", myPubKey, signa)

	fmt.Println("Is ok =>", isOk)

	return isOk

}

func DilithiumProcess() bool {

	myPubKey, myPrivateKey := pqc.GenerateDilithiumKeypair()

	fmt.Println("PubKey is ", myPubKey)

	fmt.Println("PrivateKey is ", myPrivateKey)

	signa := pqc.GenerateDilithiumSignature(myPrivateKey, "Hello KLY")

	fmt.Println("Signa is ", signa)

	isOk := pqc.VerifyDilithiumSignature("Hello KLY", myPubKey, signa)

	fmt.Println("Is ok =>", isOk)

	return isOk

}

func BlsProcess() bool {

	// Generate keypair

	privateKey, publicKey := bls.GenerateKeypair()

	fmt.Println("Privatekey is => ", privateKey)

	fmt.Println("Publickey is => ", publicKey)

	// Generate signature

	message := "Hello KLY"

	signa := bls.GenerateSignature(privateKey, message)

	fmt.Println("Signa is => ", signa)

	// Now verify (True Positive)
	fmt.Println("Is ok with norm message => ", bls.VerifySignature(publicKey, message, signa))

	// Now verify with wrong msg (True Negative)
	fmt.Println("Is ok with norm message => ", bls.VerifySignature(publicKey, "Hello badass", signa))

	// Now generate more keypairs to test aggregation

	privateKey1, publicKey1 := bls.GenerateKeypair()
	_, publicKey2 := bls.GenerateKeypair()
	_, publicKey3 := bls.GenerateKeypair()

	signa1 := bls.GenerateSignature(privateKey1, message)
	// signa2 := crypto_primitives.GenerateBlsSignature(privateKey2, message)
	// signa3 := crypto_primitives.GenerateBlsSignature(privateKey3, message)

	aggregatedSigna := bls.AggregateSignatures([]string{signa, signa1})

	fmt.Println("Aggregated signa is => ", aggregatedSigna)

	// Aggregate pubkeys

	rootPubKey := bls.AggregatePubKeys([]string{publicKey, publicKey1, publicKey2, publicKey3})

	fmt.Println("RootPubKey is => ", rootPubKey)

	// Verify with threshold

	aggregatedPubOfSigners := bls.AggregatePubKeys([]string{publicKey, publicKey1})

	fmt.Println("Aggregated 0 and 1 is => ", aggregatedPubOfSigners)

	aggregatedPub23 := bls.AggregatePubKeys([]string{publicKey2, publicKey3})

	fmt.Println("Aggregated 2 and 3 is => ", aggregatedPub23)

	fmt.Println("Their sum => ", bls.AggregatePubKeys([]string{aggregatedPubOfSigners, aggregatedPub23}))

	fmt.Println("Is threshold reached => ", bls.VerifyThresholdSignature(aggregatedPubOfSigners, aggregatedSigna, rootPubKey, message, []string{publicKey2, publicKey3}, 2))

	return true

}

type Web1337 struct {
	symbioteID, nodeURL, proxyURL string

	workflowVersion int64
}
