package tests

import (
	"fmt"
	"testing"

	kly_tbls "github.com/KLYN74R/Web1337Golang/crypto_primitives/tbls"
)

func TestTBlsProcess(t *testing.T) {

	/*

		T = 2
		N = 3

	*/

	randomIDs := kly_tbls.GenerateRandomIds(3)

	fmt.Println("IDs are => ", randomIDs)

	// Each group member do it individually

	vvec1, secretShares1 := kly_tbls.GenerateTbls(2, randomIDs)
	vvec2, secretShares2 := kly_tbls.GenerateTbls(2, randomIDs)
	vvec3, secretShares3 := kly_tbls.GenerateTbls(2, randomIDs)

	fmt.Println("Vvec 1 ", vvec1)

	// Now derive rootPubKey

	rootPubKey := kly_tbls.DeriveRootPubKey(vvec1, vvec2, vvec3)

	fmt.Println("RootPubKey is => ", rootPubKey)

	// Now imagine that members 1 and 2 aggree to sign something while member 3 - disagree. Generate partial signatures 1 and 2

	msg := "Hello World"

	secretSharesFor1 := []string{secretShares1[0], secretShares2[0], secretShares3[0]}
	secretSharesFor2 := []string{secretShares1[1], secretShares2[1], secretShares3[1]}

	partialSignature1 := kly_tbls.GeneratePartialSignature(randomIDs[0], msg, secretSharesFor1)
	partialSignature2 := kly_tbls.GeneratePartialSignature(randomIDs[1], msg, secretSharesFor2)

	fmt.Println("Partial signature 1 is => ", partialSignature1)
	fmt.Println("Partial signature 2 is => ", partialSignature2)

	// Aggregate them

	rootSignature := kly_tbls.BuildRootSignature([]string{partialSignature1, partialSignature2}, []string{randomIDs[0], randomIDs[1]})

	fmt.Println("Root signature is => ", rootSignature)

	fmt.Println("Is root signature ok ? => ", kly_tbls.VerifyRootSignature(rootPubKey, rootSignature, msg))

}
