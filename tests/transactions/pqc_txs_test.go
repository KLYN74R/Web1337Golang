package tests

// import (
// 	"encoding/json"
// 	"fmt"
// 	"testing"

// 	web1337 "github.com/KLYN74R/Web1337Golang"
// 	pqc "github.com/KLYN74R/Web1337Golang/crypto_primitives/pqc"
// )

// func PqcTransaction(t *testing.T) {

// 	myKeypairBliss := pqc.GenerateBlissKeypair()

// 	myKeypairDilithium := pqc.GenerateBlissKeypair()

// 	var (
// 		shardID              = "2VEzwUdvSRuv1k2JaAEaMiL7LLNDTUf9bXSapqccCcSb"
// 		recipient            = "nXSYHp74u88zKPiRi7t22nv4WCBHXUBpGrVw3V93f2s"
// 		from                 = myKeypair.Pub
// 		myPrivateKey         = myKeypair.Prv
// 		nonce                = uint(0)
// 		fee          float32 = 0.005
// 		amountInKLY  float32 = 0.2
// 	)

// 	myOptions := web1337.Options{
// 		SymbioteID:      "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
// 		WorkflowVersion: 1,
// 	}

// 	sdkHandler, _ := web1337.NewWeb1337(myOptions)

// 	defTx, _ := sdkHandler.CreateDefaultTransaction(shardID, from, myPrivateKey, nonce, fee, recipient, amountInKLY, nil)

// 	jsonData, _ := json.MarshalIndent(defTx, "", "  ")

// 	fmt.Println(string(jsonData))

// 	// Now create Ed25519 to Multisig tx

// 	msigRecipient := "7GPupbq1vtKUgaqVeHiDbEJcxS7sSjwPnbht4eRaDBAEJv8ZKHNCSu2Am3CuWnHjta"

// 	var rev_t int = 0

// 	msigTx, _ := sdkHandler.CreateDefaultTransaction(shardID, from, myPrivateKey, nonce, fee, msigRecipient, amountInKLY, &rev_t)

// 	jsonDataMsig, _ := json.MarshalIndent(msigTx, "", "  ")

// 	fmt.Println(string(jsonDataMsig))

// }
