package tests

import (
	"fmt"
	"testing"

	ed25519 "github.com/KLYN74R/Web1337Golang/crypto_primitives/ed25519"
)

func TestEd25519WithBips(t *testing.T) {

	keypair := ed25519.GenerateKeyPair("", "", []uint32{})

	fmt.Println("Ed25519 full => ", keypair)

	pub := "DR25VGwjEdbt2qMzD8bDEvS9gEMewvY1a5g2ZEFTimzE"
	//	prv := "MC4CAQAwBQYDK2VwBCIEIK7uW1aKknXSemukkKgdVWQJJ1BgyBYVg0/IHPV3Cb9K"

	// signa := ed25519.GenerateSignature(prv, "Hello")

	// fmt.Println("Signa => ", signa)

	fmt.Println("Is ok => ", ed25519.VerifySignature("Hello", pub, "MklxnbCmKsIsTrHinTzuvxSkQA5DHq1UwNuQ+G+1JO83Cfem1HEO8IwAffcvo8dTV5XZLOXDZshejXuZ4ca1Dw=="))

	// Check speed

	for i := 0; i < 20_000; i++ {

		ed25519.VerifySignature("Hello", pub, "MklxnbCmKsIsTrHinTzuvxSkQA5DHq1UwNuQ+G+1JO83Cfem1HEO8IwAffcvo8dTV5XZLOXDZshejXuZ4ca1Dw==")

	}

}
