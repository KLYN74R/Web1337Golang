package crypto_primitives

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
	"go.dedis.ch/kyber/v4/pairing/bn256"
	bls "go.dedis.ch/kyber/v4/sign/bls"
	"go.dedis.ch/kyber/v4/util/random"
)

func GenerateBlsKeypair() (privateKeyAsHex, publicKeyAsBase58 string) {

	suite := bn256.NewSuiteG2()

	privateKey, publicKey := bls.NewKeyPair(suite, random.New())

	pubKeyToExport, _ := publicKey.MarshalBinary()

	return privateKey.String(), base58.Encode(pubKeyToExport)

}

func AggregateBlsPubKeys(arrayOfPubKeysAsBase58 []string) (aggregatedPubKeyAsBase58 string) {

	suite := bn256.NewSuite()

	finalPubKey := suite.G2().Point()

	for _, pubKeyAsBase58 := range arrayOfPubKeysAsBase58 {

		pubKeyAsBytes := base58.Decode(pubKeyAsBase58)

		recoveredPubKey := suite.G2().Point()

		err := recoveredPubKey.UnmarshalBinary(pubKeyAsBytes)

		if err != nil {

			panic("Can't recover pubkey")

		} else {

			finalPubKey.Add(finalPubKey, recoveredPubKey)

		}

	}

	// Now encode final pubkey to Base58 and return

	finalPubKeyToExport, _ := finalPubKey.MarshalBinary()

	return base58.Encode(finalPubKeyToExport)

}

func AggregateBlsSignatures(arrayOfSignaturesAsBase64 []string) string {

	suite := bn256.NewSuite()

	// Run a cycle to convert signatures from Base64 to []byte

	// signaturesSet := make([]([]byte), len(arrayOfSignaturesAsBase64))

	var signaturesSet [][]byte

	for _, signatureAsBase64 := range arrayOfSignaturesAsBase64 {

		signatureAsBytes, err := base64.StdEncoding.DecodeString(signatureAsBase64)

		if err != nil {

			panic(err)

		}

		signaturesSet = append(signaturesSet, signatureAsBytes)

	}

	// Now pass signatures to BLS package and then return the Base64 encoded aggregated signature

	aggregatedSignature, _ := bls.AggregateSignatures(suite, signaturesSet...)

	return base64.StdEncoding.EncodeToString(aggregatedSignature)

}

func GenerateBlsSignature(privateKeyAsHex, message string) (blsSignature string) {

	msg := []byte(message)

	suite := bn256.NewSuite()

	// Recover private key

	privateKeyAsBytes, _ := hex.DecodeString(privateKeyAsHex)

	recoveredPrivateKey := suite.G2().Scalar()

	err := recoveredPrivateKey.UnmarshalBinary(privateKeyAsBytes)

	if err != nil {

		panic("Impossible to recover private key")

	}

	blsSignaAsBytes, _ := bls.Sign(suite, recoveredPrivateKey, msg)

	// To base64

	return base64.StdEncoding.EncodeToString(blsSignaAsBytes)

}

func VerifyBlsSignature(pubKeyAsBase58, message, signatureAsBase64 string) bool {

	msg := []byte(message)

	suite := bn256.NewSuite()

	// Recover public key

	pubKeyAsBytes := base58.Decode(pubKeyAsBase58)

	recoveredPubKey := suite.G2().Point()

	err := recoveredPubKey.UnmarshalBinary(pubKeyAsBytes)

	if err != nil {

		panic("Impossible to recover public key")

	}

	// Decode signature from Base64

	signatureAsBytes, _ := base64.StdEncoding.DecodeString(signatureAsBase64)

	err = bls.Verify(suite, recoveredPubKey, msg, signatureAsBytes)

	return err == nil

}

func VerifyBlsThresholdSignature(aggregatedPubkeyWhoSignAsBase58, aggregatedSignatureAsBase64, rootPubAsBase58, message string, afkPubkeysArray []string, reverseThreshold uint) bool {

	if len(afkPubkeysArray) <= int(reverseThreshold) {

		verifiedSignature := VerifyBlsSignature(

			aggregatedPubkeyWhoSignAsBase58,
			message,
			aggregatedSignatureAsBase64,
		)

		if verifiedSignature {

			// If all the previos steps are OK - do the most CPU intensive task - pubkeys aggregation

			suite := bn256.NewSuite()

			// Recover public key of parts who signed this message and root pubkey

			aggregatedPubKeyAsBytes := base58.Decode(aggregatedPubkeyWhoSignAsBase58)
			rootPubKeyAsBytes := base58.Decode(aggregatedPubkeyWhoSignAsBase58)

			// Create empty templates

			recoveredPubKeyOfSigners := suite.G2().Point()
			recoveredRootPubKey := suite.G2().Point()

			err1 := recoveredPubKeyOfSigners.UnmarshalBinary(aggregatedPubKeyAsBytes)
			err2 := recoveredRootPubKey.UnmarshalBinary(rootPubKeyAsBytes)

			if err1 != nil || err2 != nil {

				return false

			}

			//Now aggregate AFK signers

			aggregatedPubKeyOfAfkSigners := AggregateBlsPubKeys(afkPubkeysArray)

			return AggregateBlsPubKeys([]string{aggregatedPubKeyOfAfkSigners, aggregatedPubkeyWhoSignAsBase58}) == rootPubAsBase58

		} else {

			return false

		}

	} else {

		return false

	}

}
