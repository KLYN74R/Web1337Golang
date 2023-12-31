package crypto_primitives

import (
	bls "github.com/herumi/bls-eth-go-binary/bls"
)

// Function to generate random IDs for N parties of DKG
func GenerateTblsRandomIDs(numberOfParties uint) (idsAsHex []string) {

	bls.Init(bls.BLS12_381)

	for i := 0; i < int(numberOfParties); i++ {

		secretkey := bls.SecretKey{}

		secretkey.SetByCSPRNG()

		idsAsHex = append(idsAsHex, secretkey.SerializeToHexStr())

	}

	return

}

// Function to generate verification vector and secret shares to distribute them among N parties
func GenerateTbls(threshold uint, yourIdAsHex string, partiesIDs []string) (verificationVector []bls.PublicKey, sharesForOtherParties []bls.SecretKey) {

	bls.Init(bls.BLS12_381)

	var secretVector []bls.SecretKey

	// Generate a secretVector and verificationVector

	for i := 0; i < int(threshold); i++ {

		secretkey := bls.SecretKey{}

		secretkey.SetByCSPRNG()

		secretVector = append(secretVector, secretkey)

		publicKey := secretkey.GetPublicKey()

		verificationVector = append(verificationVector, *publicKey)

	}

	// Generate key shares for other parties

	for _, hexID := range partiesIDs {

		sk := bls.SecretKey{}

		sk.SetByCSPRNG()

		emptyID := &bls.ID{}

		emptyID.DeserializeHexStr(hexID)

		sk.Set(secretVector, emptyID)

		sharesForOtherParties = append(sharesForOtherParties, sk)

	}

	return

}

// Function to verify share which you'll get from other parties based on their verification vector + your ID
func VerifyTblsShare(yourIdAsHex, secretShareAsHex string, verificationVectorAsHex []string) bool {

	bls.Init(bls.BLS12_381)

	pubKeyToVerifyShare := bls.PublicKey{}

	yourID := &bls.ID{}

	yourID.DeserializeHexStr(yourIdAsHex)

	// Deserialize verification vector ( []Hex => []bls.PublicKey )

	vvec := make([]bls.PublicKey, len(verificationVectorAsHex))

	for index, hexValueInVerificationVector := range verificationVectorAsHex {

		vvec[index].DeserializeHexStr(hexValueInVerificationVector)

	}

	pubKeyToVerifyShare.Set(vvec, yourID)

	// Deserialize secret share that you get from other signer

	secretShare := bls.SecretKey{}

	secretShare.DeserializeHexStr(secretShareAsHex)

	pubKeyFromShare := *secretShare.GetPublicKey()

	return pubKeyFromShare.IsEqual(&pubKeyToVerifyShare)

}
