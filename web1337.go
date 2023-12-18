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

type Web1337 struct {
}
