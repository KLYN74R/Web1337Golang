package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	web1337 "github.com/KLYN74R/Web1337Golang"
	kly_bls "github.com/KLYN74R/Web1337Golang/crypto_primitives/bls"
)

func TestBlsTransactions(t *testing.T) {

	_, publicKey1 := kly_bls.GenerateKeypair()
	_, publicKey2 := kly_bls.GenerateKeypair()
	_, publicKey3 := kly_bls.GenerateKeypair()

	rootPubKey := kly_bls.AggregatePubKeys([]string{publicKey1, publicKey2, publicKey3})

	aggregatedPubOfActive := kly_bls.AggregatePubKeys([]string{publicKey1, publicKey2})

	var (
		// shardID      = "2VEzwUdvSRuv1k2JaAEaMiL7LLNDTUf9bXSapqccCcSb"
		recipient           = "nXSYHp74u88zKPiRi7t22nv4WCBHXUBpGrVw3V93f2s"
		from                = rootPubKey
		nonce       uint    = 0
		fee         float32 = 0.00005
		amountInKLY float32 = 0.2

		afkSigners = []string{publicKey3}
	)

	myOptions := web1337.Options{
		SymbioteID:      "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	defTx := sdkHandler.CreateMultisigTransaction(from, aggregatedPubOfActive, "", afkSigners, nonce, fee, recipient, amountInKLY, nil)

	jsonData, _ := json.MarshalIndent(defTx, "", " ")

	fmt.Println(string(jsonData))

	// let signedTx = await web1337.createDefaultTransaction(shardID,from,myPrivateKey,nonce,recipient,fee,amountInKLY);

	// console.log(signedTx);

}
