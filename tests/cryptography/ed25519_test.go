package tests

import (
	"fmt"
	"testing"

	ed25519 "github.com/KLYN74R/Web1337Golang/crypto_primitives/ed25519"
)

func TestEd25519(t *testing.T) {

	mnemonic := "smoke suggest security index situate almost ethics tone wash crystal debris mosquito pony extra husband elder over relax width occur inspire keen sudden average"

	mnemonicPassword := ""

	bip44Path := []uint32{44, 7331, 0, 0}

	keypair := ed25519.GenerateKeyPair(mnemonic, mnemonicPassword, bip44Path)

	fmt.Println(keypair)

	// signa := ed25519.GenerateSignature(keypair.Prv, "BlaBlaBla")

	// fmt.Println("Signa => ", signa)

	// fmt.Println("Is ok => ", ed25519.VerifySignature("BlaBlaBla", keypair.Pub, signa))

}
