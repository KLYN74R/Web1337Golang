package web1337

func (sdk *Web1337) GetContractMetadata(contractId string) ([]byte, error) {
	return sdk.getRequest("/account/" + contractId)
}

func (sdk *Web1337) GetContractStorage(contractId, storageName string) ([]byte, error) {
	return sdk.getRequest("/account/" + contractId + "_STORAGE_" + storageName)
}

func (web1337 *Web1337) CallContract(originShard, yourPub, yourPrv, contractId, method string, params map[string]interface{}, injects []string) {

	/*

	   Payload is

	   {

	       contractID:<BLAKE3 hashID of contract OR alias of contract(for example, system contracts)>,
	       method:<string method to call>,
	       gasLimit:<maximum allowed in KLY to execute contract>
	       params:[] params to pass to function
	       imports:[] imports which should be included to contract instance to call. Example ['default.CROSS-CONTRACT','storage.GET_FROM_ARWEAVE']. As you understand, it's form like <MODULE_NAME>.<METHOD_TO_IMPORT>

	   }


	*/

	// coreWorkflowVersion := web1337.Symbiotes[web1337.CurrentSymbiote].WorkflowVersion

	// payload := map[string]interface{}{
	// 	"type":   SIGNATURES_TYPES.DEFAULT_SIG,
	// 	"to":     recipient,
	// 	"amount": amountInKLY,
	// }

	// // In case we send from Ed25519 to BLS
	// if rev_t != nil {
	// 	payload["rev_t"] = *rev_t
	// }

	// txTemplate := web1337.GetTransactionTemplate(coreWorkflowVersion, yourAddress, TXS_TYPES.TX, nonce, fee, payload)

	// return sdk.getRequest("/symbiote_info")
}

func (sdk *Web1337) deployContractToWvm(bytecode string) {

	// Payload is

	// {
	// 	bytecode:<hexString>,
	// 	lang:<RUST|ASC>,
	// 	constructorParams:[]
	// }

	// return sdk.getRequest("/quorum_thread_checkpoint")
}

func (sdk *Web1337) subscribeForEventsByContract(contractId, eventId string) {
	// return sdk.getRequest("/quorum_thread_checkpoint")
}

// func (web1337 *Web1337) CreateDefaultTransaction(originShard, yourAddress, yourPrivateKeyAsHex string, nonce uint, fee float32, recipient string, amountInKLY float32, rev_t *int) (TransactionTemplate, error) {

// 	coreWorkflowVersion := web1337.Symbiotes[web1337.CurrentSymbiote].WorkflowVersion

// 	payload := map[string]interface{}{
// 		"type":   SIGNATURES_TYPES.DEFAULT_SIG,
// 		"to":     recipient,
// 		"amount": amountInKLY,
// 	}

// 	// In case we send from Ed25519 to BLS
// 	if rev_t != nil {
// 		payload["rev_t"] = *rev_t
// 	}

// 	txTemplate := web1337.GetTransactionTemplate(coreWorkflowVersion, yourAddress, TXS_TYPES.TX, nonce, fee, payload)

// 	dataToSign := web1337.CurrentSymbiote + string(coreWorkflowVersion) + originShard + TXS_TYPES.TX + fmt.Sprintf("%v", payload) + string(nonce) + fmt.Sprintf("%f", fee)

// 	txTemplate.Sig = ed25519.GenerateSignature(yourPrivateKeyAsHex, dataToSign)

// 	return txTemplate, nil
// }
