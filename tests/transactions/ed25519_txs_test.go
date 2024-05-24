package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	web1337 "github.com/KLYN74R/Web1337Golang"
	ed25519 "github.com/KLYN74R/Web1337Golang/crypto_primitives/ed25519"
)

func TestEd25519Transactions(t *testing.T) {

	myKeypair := ed25519.Ed25519Box{

		Mnemonic:  "smoke suggest security index situate almost ethics tone wash crystal debris mosquito pony extra husband elder over relax width occur inspire keen sudden average",
		Bip44Path: []uint32{44, 7331, 0, 0},
		Pub:       "8LZ1XWiLHwuEVsWfaWmY2sm98hYrFAbGqfi4zwbApfJT",
		Prv:       "MC4CAQAwBQYDK2VwBCIEIHmKXvMPju2bdhbhqUYR08ujs/SYH/WPC3cPQfhnym89",
	}

	var (
		shardID      = "2VEzwUdvSRuv1k2JaAEaMiL7LLNDTUf9bXSapqccCcSb"
		recipient    = "nXSYHp74u88zKPiRi7t22nv4WCBHXUBpGrVw3V93f2s"
		from         = myKeypair.Pub
		myPrivateKey = myKeypair.Prv
		nonce        = 0
		fee          = 1
		amountInKLY  = 13
	)

	myOptions := web1337.Options{
		SymbioteID:      "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	defTx, _ := sdkHandler.CreateDefaultTransaction(shardID, from, myPrivateKey, nonce, fee, recipient, amountInKLY, nil)

	jsonData, _ := json.Marshal(defTx)

	fmt.Println(string(jsonData))

	// let signedTx = await web1337.createDefaultTransaction(shardID,from,myPrivateKey,nonce,recipient,fee,amountInKLY);

	// console.log(signedTx);

}
