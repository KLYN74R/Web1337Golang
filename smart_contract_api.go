package web1337

import (
	"fmt"

	ed25519 "github.com/KLYN74R/Web1337Golang/crypto_primitives/ed25519"
	pqc "github.com/KLYN74R/Web1337Golang/crypto_primitives/pqc"
	SIGNATURES_TYPES "github.com/KLYN74R/Web1337Golang/signatures_types"
	TXS_TYPES "github.com/KLYN74R/Web1337Golang/txs_types"
)

func (sdk *Web1337) GetContractMetadata(contractId string) ([]byte, error) {
	return sdk.getRequest("/account/" + contractId)
}

func (sdk *Web1337) GetContractStorage(contractId, storageName string) ([]byte, error) {
	return sdk.getRequest("/account/" + contractId + "_STORAGE_" + storageName)
}

func (sdk *Web1337) deployContractToWvm(web1337 *Web1337, originShard, yourAddress, yourPrivateKey, sigType, bytecode, lang string, nonce uint, fee float32, constructorParams []string) TransactionTemplate {

	/*

	   Full transaction which contains contract deploy must have such structure

	   {
	       v: 0,
	       creator: '2VEzwUdvSRuv1k2JaAEaMiL7LLNDTUf9bXSapqccCcSb',
	       type: 'CONTRACT_DEPLOY',
	       nonce: 0,
	       fee: 1,
	       payload: {
	           type: 'D',
	           bytecode:<hexString>,
	           lang:<RUST|ASC>,
	           constructorParams:[]
	       },
	       sig: '5AGkLlK3knzYZeZwjHKPzlX25lPMd7nU+rR5XG9RZa3sDpYrYpfnzqecm5nNONnl5wDcxmjOkKMbO7ulAwTFDQ=='
	   }

	*/

	workflowVersion := web1337.Symbiotes[web1337.CurrentSymbiote].WorkflowVersion

	payload := map[string]interface{}{
		"type":              sigType,
		"bytecode":          bytecode,
		"lang":              lang,
		"constructorParams": constructorParams,
	}

	txTemplate := sdk.GetTransactionTemplate(workflowVersion, yourAddress, TXS_TYPES.CONTRACT_DEPLOY, nonce, fee, payload)

	dataToSign := fmt.Sprintf("%s%d%s%s%s%d%f", web1337.CurrentSymbiote, workflowVersion, originShard, TXS_TYPES.CONTRACT_DEPLOY, mapToJSON(payload), nonce, fee)

	switch sigType {

	case SIGNATURES_TYPES.DEFAULT_SIG:
		txTemplate.Sig = ed25519.GenerateSignature(yourPrivateKey, dataToSign)

	case SIGNATURES_TYPES.POST_QUANTUM_BLISS:
		txTemplate.Sig = pqc.GenerateBlissSignature(yourPrivateKey, dataToSign)

	case SIGNATURES_TYPES.POST_QUANTUM_DIL:
		txTemplate.Sig = pqc.GenerateDilithiumSignature(yourPrivateKey, dataToSign)

	}

	return txTemplate

}

func (web1337 *Web1337) CallContract(originShard, yourPub, yourPrv, contractId, method string, params map[string]interface{}, injects []string) {

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

func (sdk *Web1337) subscribeForEventsByContract(contractId, eventId string) {
	// return sdk.getRequest("/quorum_thread_checkpoint")
}
