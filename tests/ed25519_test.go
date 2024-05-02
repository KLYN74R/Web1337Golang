package tests

import (
	"fmt"
	"testing"

	ed25519 "github.com/KLYN74R/Web1337Golang/crypto_primitives/ed25519"
)

func TestEd25519(t *testing.T) {

	keypair := ed25519.GenerateKeyPair("smoke suggest security index situate almost ethics tone wash crystal debris mosquito pony extra husband elder over relax width occur inspire keen sudden average", "", []uint32{})

	fmt.Println("Ed25519 full => ", keypair)

	signa := ed25519.GenerateSignature(keypair.Prv, "BlaBlaBla")

	fmt.Println("Signa => ", signa)

	fmt.Println("Is ok => ", ed25519.VerifySignature("BlaBlaBla", keypair.Pub, signa))

}
