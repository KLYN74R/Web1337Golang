/*

Web1337 by KLY

For Golang devs

*/

package web1337

import (
	"fmt"

	"github.com/KLYN74R/Web1337Golang/crypto_primitives"
)

func Ed25519Process() bool {

	myPubKey, myPrivateKey := crypto_primitives.GenerateEd25519KeyPair()

	fmt.Println("PubKey is ", myPubKey)

	fmt.Println("PrivateKey is ", myPrivateKey)

	signa := crypto_primitives.GenerateEd25519Signature(myPrivateKey, "Hello KLY")

	fmt.Println("Signa is ", signa)

	isOk := crypto_primitives.VerifyEd25519Signature("Hello KLY", myPubKey, signa)

	fmt.Println("Is ok =>", isOk)

	return isOk

}

func BlissProcess() bool {

	myPubKey, myPrivateKey := crypto_primitives.GenerateBlissKeypair()

	fmt.Println("PubKey is ", myPubKey)

	fmt.Println("PrivateKey is ", myPrivateKey)

	signa := crypto_primitives.GenerateBlissSignature(myPrivateKey, "Hello KLY")

	fmt.Println("Signa is ", signa)

	isOk := crypto_primitives.VerifyBlissSignature("Hello KLY", myPubKey, signa)

	fmt.Println("Is ok =>", isOk)

	return isOk

}

func DilithiumProcess() bool {

	myPubKey, myPrivateKey := crypto_primitives.GenerateDilithiumKeypair()

	fmt.Println("PubKey is ", myPubKey)

	fmt.Println("PrivateKey is ", myPrivateKey)

	signa := crypto_primitives.GenerateDilithiumSignature(myPrivateKey, "Hello KLY")

	fmt.Println("Signa is ", signa)

	isOk := crypto_primitives.VerifyDilithiumSignature("Hello KLY", myPubKey, signa)

	fmt.Println("Is ok =>", isOk)

	return isOk

}

func BlsProcess() bool {

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

	return true

}

type Web1337 struct {
	symbioteID, nodeURL, proxyURL string

	workflowVersion int64
}
