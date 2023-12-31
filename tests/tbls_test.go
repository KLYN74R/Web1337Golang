package tests

import (
	"fmt"

	"testing"

	bls "github.com/herumi/bls-eth-go-binary/bls"
)

func TestGenerateContribution(t *testing.T) {

	//_____________generateContribution_____________

	bls.Init(bls.BLS12_381)

	ids := []string{

		"4bf5122f344554c53bde2ebb8cd2b7e3d1600ad631c385a5d7cce23c7785451a",
		"0000000000000000000000000000000000000000000000000000000000000002",
		"0000000000000000000000000000000000000000000000000000000000000003"} // we'll try 2/3 threshold

	//this id's verification vector
	var vvec1 []bls.PublicKey

	// this id's secret keys
	var svec1 []bls.SecretKey

	// this id's sk contributions shares
	var skContribution1 []bls.SecretKey

	// generate a sk and vvec

	for i := 0; i < 2; i++ {

		secretkey := bls.SecretKey{}

		secretkey.SetByCSPRNG()

		svec1 = append(svec1, secretkey)

		publicKey := secretkey.GetPublicKey()

		vvec1 = append(vvec1, *publicKey)

	}

	// generate key shares

	for _, hexID := range ids {

		sk := bls.SecretKey{}

		sk.SetByCSPRNG()

		emptyID := &bls.ID{}

		// fmt.Println("Empty ID is bla => ", emptyID.SerializeToHexStr())

		emptyID.DeserializeHexStr(hexID)

		sk.Set(svec1, emptyID)

		skContribution1 = append(skContribution1, sk)

	}

	fmt.Println("Finally")
	fmt.Println("Verification vector 1 => ", vvec1)
	fmt.Println("Secret key contribution 1 => ", skContribution1)

	// return vvec, skContribution

	//___________________verifyContributionShare____________________

	// Since we have index 1 - verify appropraite share

	pubKeyToVerifyShare := bls.PublicKey{}

	idWithIndexOne := &bls.ID{}

	idWithIndexOne.DeserializeHexStr(ids[0])

	pubKeyToVerifyShare.Set(vvec1, idWithIndexOne)

	pubKeyFromShare := *skContribution1[0].GetPublicKey()

	fmt.Println("Is share OK ? => ", pubKeyFromShare.IsEqual(&pubKeyToVerifyShare))

	// Repeat generation of verification vector for participants 2 and 3

	var vvec2 []bls.PublicKey
	var vvec3 []bls.PublicKey

	var svec2 []bls.SecretKey
	var svec3 []bls.SecretKey

	var skContribution2 []bls.SecretKey
	var skContribution3 []bls.SecretKey

	// generate a sk and vvec

	for i := 0; i < 2; i++ {

		secretkey2 := bls.SecretKey{}

		secretkey2.SetByCSPRNG()

		secretkey3 := bls.SecretKey{}

		secretkey3.SetByCSPRNG()

		svec2 = append(svec2, secretkey2)
		svec3 = append(svec3, secretkey3)

		publicKey2 := secretkey2.GetPublicKey()
		publicKey3 := secretkey3.GetPublicKey()

		vvec2 = append(vvec2, *publicKey2)
		vvec3 = append(vvec3, *publicKey3)

	}

	// generate key shares

	for _, hexID := range ids {

		sk2 := bls.SecretKey{}

		sk2.SetByCSPRNG()

		sk3 := bls.SecretKey{}

		sk3.SetByCSPRNG()

		emptyID2 := &bls.ID{}
		emptyID3 := &bls.ID{}

		emptyID2.DeserializeHexStr(hexID)
		emptyID3.DeserializeHexStr(hexID)

		sk2.Set(svec2, emptyID2)
		sk3.Set(svec3, emptyID3)

		skContribution2 = append(skContribution2, sk2)
		skContribution3 = append(skContribution3, sk3)

	}

	// Try to get the master pubkey
	// Based on verification vectors

	fmt.Println("\n\n===========================")

	// skShares1 := []bls.SecretKey{skContribution1[0], skContribution2[0], skContribution3[0]}
	// skShares2 := []bls.SecretKey{skContribution1[1], skContribution2[1], skContribution3[1]}
	// skShares3 := []bls.SecretKey{skContribution1[2], skContribution2[2], skContribution3[2]}

	// rootPub1 := bls.GetMasterPublicKey(skShares1)
	// rootPub2 := bls.GetMasterPublicKey(skShares2)
	// rootPub3 := bls.GetMasterPublicKey(skShares3)

	// fmt.Println("Len 1 => ", rootPub1[0].SerializeToHexStr())
	// fmt.Println("Len 2 => ", rootPub2[0].SerializeToHexStr())
	// fmt.Println("Len 3 => ", rootPub3[0].SerializeToHexStr())

	allVerificationVectors := [][]bls.PublicKey{vvec2, vvec3}

	rootPubKey := []bls.PublicKey{vvec1[0], vvec1[1]}

	for _, singleVerificationVector := range allVerificationVectors {

		for index, pubKey := range singleVerificationVector {

			rootPubKey[index].Add(&pubKey)

			// if rootPubKey[index] == nil {}

		}

	}

	// fmt.Println("RootPubKey is => ", rootPubKey[0].SerializeToHexStr())

	finalRootPub := rootPubKey[0]

	//____________________ Generate partial signatures __________________
	// Saying 1 and 2 sign message, 3 is AFK
	// dkg.addContributionShares

	message := "hello world"

	emptyID1 := &bls.ID{}
	emptyID2 := &bls.ID{}

	emptyID1.DeserializeHexStr(ids[0])
	emptyID2.DeserializeHexStr(ids[1])

	skShares1 := []bls.SecretKey{skContribution1[0], skContribution2[0], skContribution3[0]}
	skShares2 := []bls.SecretKey{skContribution1[1], skContribution2[1], skContribution3[1]}
	// skShares3 := []bls.SecretKey{skContribution1[2], skContribution2[2], skContribution3[2]}

	skShares1[0].Add(&skShares1[1])
	skShares1[0].Add(&skShares1[2])

	skShares2[0].Add(&skShares2[1])
	skShares2[0].Add(&skShares2[2])

	groupSecretReceivedBy1 := skShares1[0]
	groupSecretReceivedBy2 := skShares2[0]

	partialSignatureBy1 := groupSecretReceivedBy1.Sign(message)
	partialSignatureBy2 := groupSecretReceivedBy2.Sign(message)

	fmt.Println("Partial signature 1 is ", partialSignatureBy1)
	fmt.Println("Partial signature 2 is ", partialSignatureBy2)

	// Build the master signature with these partial ones

	rootSignature := bls.Sign{}

	rootSignature.Recover([]bls.Sign{*partialSignatureBy1, *partialSignatureBy2}, []bls.ID{*emptyID1, *emptyID2})

	fmt.Println("Root signature is => ", rootSignature)

	fmt.Println("Is root signature ok ? => ", rootSignature.Verify(&finalRootPub, message))

	fmt.Println("Root pubkey is => ", finalRootPub.SerializeToHexStr())
	fmt.Println("Root signature is => ", rootSignature.SerializeToHexStr())

}
