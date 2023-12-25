package crypto_primitives

import (
	"strings"

	"github.com/coinbase/kryptology/pkg/ted25519/ted25519"

	"encoding/hex"
)

// Function to generate pubkey, shares for signers and commitments
func GenerateTed25519(T, N int) (myPubKey string, sharesToExport, commitmentsToExport []string) {

	//Create T/N threshold configs
	config := ted25519.ShareConfiguration{T, N}

	// GenerateSharedKey generates a random key, splits it, and returns the public key, shares, and VSS commitments.
	// func GenerateSharedKey(config *ShareConfiguration) (PublicKey, []*KeyShare, Commitments, error)

	myPub, secretShares, commitments, _ := ted25519.GenerateSharedKey(&config)

	//Serialize secret shares and commitments for VSS

	for _, singleShare := range secretShares {

		sharesToExport = append(sharesToExport, hex.EncodeToString(singleShare.Bytes()))

	}

	for _, commitmentProof := range commitments.CommitmentsToBytes() {

		//Per user
		commitmentsToExport = append(commitmentsToExport, hex.EncodeToString(commitmentProof))

	}

	return hex.EncodeToString(myPub.Bytes()), sharesToExport, commitmentsToExport

}

// Function to verify share received from one of signers based on this share and array of commitments
func VerifySecretShareTed25519(T, N int, receivedSecretShareAsHex string, receivedCommitmentsAsHex []string) bool {

	//Create T/N threshold configs
	config := ted25519.ShareConfiguration{T, N}

	//Deserialize my part of secret shares
	mySecretShareBuffer, _ := hex.DecodeString(receivedSecretShareAsHex)

	commitments := make([][]byte, T)

	for i := range commitments {

		commitments[i] = make([]byte, 32)

	}

	for i, singleCommitmentAsHex := range receivedCommitmentsAsHex {

		commitmentAsBytes, _ := hex.DecodeString(singleCommitmentAsHex)

		commitments[i] = commitmentAsBytes

	}

	finalCommitments, _ := ted25519.CommitmentsFromBytes(commitments)

	ok, _ := ted25519.KeyShareFromBytes(mySecretShareBuffer).VerifyVSS(finalCommitments, &config)

	return ok

}

func GenerateNonceSharesTed25519(T, N int, secretShareAsHex, myPubKeyAsHex, message string) (noncePubKeyAsHex string, nonceSharesAsHex, nonceCommitmentsAsHex []string) {

	config := ted25519.ShareConfiguration{T, N}

	secretShareAsBuffer, _ := hex.DecodeString(secretShareAsHex)

	rootPubKeyAsBuffer, _ := hex.DecodeString(myPubKeyAsHex)

	rootPubKey, _ := ted25519.PublicKeyFromBytes(rootPubKeyAsBuffer)

	msgAsBytes, _ := hex.DecodeString(message)

	typedMsg := ted25519.Message(msgAsBytes)

	noncePub, nonceShares, nonceCommitments, _ := ted25519.GenerateSharedNonce(&config, ted25519.KeyShareFromBytes(secretShareAsBuffer), rootPubKey, typedMsg)

	for _, singleNonceShare := range nonceShares {

		nonceSharesAsHex = append(nonceSharesAsHex, hex.EncodeToString(singleNonceShare.Bytes()))

	}

	for _, singleNonceCommitment := range nonceCommitments.CommitmentsToBytes() {

		nonceCommitmentsAsHex = append(nonceCommitmentsAsHex, hex.EncodeToString(singleNonceCommitment))

	}

	return hex.EncodeToString(noncePub.Bytes()), nonceSharesAsHex, nonceCommitmentsAsHex

}

func SubsignTed25519(secretShareAsHex, rootPubKeyAsHex, message string, nonceSharesAsHex, noncePubKeysAsHex []string) string {

	//Deserialize secret share byte buffer received by you initially(1st communications round)
	receivedSecretShareBuffer, _ := hex.DecodeString(secretShareAsHex)

	//Deserialize common(general) pubkey
	rootPubKeyBuffer, _ := hex.DecodeString(rootPubKeyAsHex)

	rootPubKey, _ := ted25519.PublicKeyFromBytes(rootPubKeyBuffer)

	messageAsBytes, _ := hex.DecodeString(message)

	//------------------------Циклы по hexNonceShares,hexNoncePubKeys для того чтоб собрать шары и публичный ключ------------------------

	var myNonceShare *ted25519.NonceShare

	//Agregate my nonceShares received from other participants
	for i, singleNonceShare := range nonceSharesAsHex {

		subBuffer, _ := hex.DecodeString(singleNonceShare) //32 bytes buffer

		if i == 0 {

			myNonceShare = ted25519.NonceShareFromBytes(subBuffer)

		} else {

			myNonceShare.Add(ted25519.NonceShareFromBytes(subBuffer))

		}

	}

	//------------------------------Agregate noncePubkeys------------------------------

	var myNoncePub ted25519.PublicKey

	for i, singleNoncePubKey := range noncePubKeysAsHex {

		subBuffer, _ := hex.DecodeString(singleNoncePubKey) // 32 bytes buffer

		subPubKey, _ := ted25519.PublicKeyFromBytes(subBuffer)

		if i == 0 {

			myNoncePub = subPubKey

		} else {

			myNoncePub = ted25519.GeAdd(myNoncePub, subPubKey)

		}

	}

	return hex.EncodeToString(ted25519.TSign(messageAsBytes, ted25519.KeyShareFromBytes(receivedSecretShareBuffer), rootPubKey, myNonceShare, myNoncePub).Bytes())

}

func AggregateSubSignaturesTed25519(T, N int, hexSubSignatures []string) string {

	//https://github.com/coinbase/kryptology/blob/269410e1b06b43da82caf28cf99cb8c0c140b65d/pkg/ted25519/ted25519/partialsig.go#L19

	//Create T/N threshold configs
	config := ted25519.ShareConfiguration{T, N}

	subSignaturesArray := make([]*ted25519.PartialSignature, T)

	for i, singleSubSig := range hexSubSignatures {

		metaToRecover := strings.Split(singleSubSig, "*")

		index, _ := hex.DecodeString(metaToRecover[0])

		subBuffer, _ := hex.DecodeString(metaToRecover[1])

		subSignaturesArray[i] = ted25519.NewPartialSignature(index[0], subBuffer)

	}

	//Build noncePub from all subPubs
	//noncePub := ted25519.GeAdd(ted25519.GeAdd(noncePub1, noncePub2), noncePub3)

	//Build full signature from subsignatures
	sig, _ := ted25519.Aggregate(subSignaturesArray, &config)

	return hex.EncodeToString(sig)

}

func VerifyTed25519(rootPubAsHex, messageAsHex, aggregatedRootSignatureAsHex string) bool {

	rootPubkeyAsBuffer, _ := hex.DecodeString(rootPubAsHex)

	rootPubKey, _ := ted25519.PublicKeyFromBytes(rootPubkeyAsBuffer)

	msg, _ := hex.DecodeString(messageAsHex)

	aggregatedSignature, _ := hex.DecodeString(aggregatedRootSignatureAsHex)

	//Check
	ok, _ := ted25519.Verify(rootPubKey, msg, aggregatedSignature)

	return ok

}
