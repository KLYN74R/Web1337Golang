package web1337

import (
	"fmt"

	bls "github.com/KLYN74R/Web1337Golang/crypto_primitives/bls"
	ed25519 "github.com/KLYN74R/Web1337Golang/crypto_primitives/ed25519"
	"github.com/KLYN74R/Web1337Golang/crypto_primitives/pqc"
	tbls "github.com/KLYN74R/Web1337Golang/crypto_primitives/tbls"

	SIGNATURES_TYPES "github.com/KLYN74R/Web1337Golang/signature_types"
)

type TransactionTemplate struct {
	V       uint                   `json:"v"`
	Creator string                 `json:"creator"`
	Type    string                 `json:"type"`
	Nonce   uint                   `json:"nonce"`
	Fee     float32                `json:"fee"`
	Payload map[string]interface{} `json:"payload"` // might be various - depends on transaction type
	SigType string                 `json:"sigType"`
	Sig     string                 `json:"sig"`
}

func (web1337 *Web1337) GetTransactionTemplate(workflowVersion uint, creator, txType, sigType string, nonce uint, fee float32, payload map[string]interface{}) TransactionTemplate {
	return TransactionTemplate{
		V:       workflowVersion,
		Creator: creator,
		Type:    txType,
		Nonce:   nonce,
		Fee:     fee,
		Payload: payload,
		SigType: sigType,
		Sig:     "",
	}
}

func (web1337 *Web1337) CreateEd25519Transaction(originShard, txType, yourAddress, base64PrivateKey string, nonce uint, fee float32, payload map[string]interface{}) (TransactionTemplate, error) {

	coreWorkflowVersion := web1337.Chains[web1337.CurrentChain].WorkflowVersion

	txTemplate := web1337.GetTransactionTemplate(coreWorkflowVersion, yourAddress, txType, SIGNATURES_TYPES.DEFAULT_SIG, nonce, fee, payload)

	dataToSign := web1337.CurrentChain + string(coreWorkflowVersion) + originShard + txType + mapToJSON(payload) + string(nonce) + fmt.Sprintf("%f", fee)

	txTemplate.Sig = ed25519.GenerateSignature(base64PrivateKey, dataToSign)

	// Return signed transaction

	return txTemplate, nil
}

func (web1337 *Web1337) SignDataForMultisigTransaction(originShard, txType, blsPrivateKey string, nonce uint, fee float32, payload map[string]interface{}) string {

	coreWorkflowVersion := web1337.Chains[web1337.CurrentChain].WorkflowVersion

	dataToSign := web1337.CurrentChain + string(coreWorkflowVersion) + originShard + txType + mapToJSON(payload) + string(nonce) + fmt.Sprintf("%f", fee)

	blsSingleSigna := bls.GenerateSignature(blsPrivateKey, dataToSign)

	return blsSingleSigna

}

func (web1337 *Web1337) CreateMultisigTransaction(txType, rootPubKey, aggregatedSignatureOfActive string, nonce uint, fee float32, payload map[string]interface{}) TransactionTemplate {

	coreWorkflowVersion := web1337.Chains[web1337.CurrentChain].WorkflowVersion

	multisigTransaction := web1337.GetTransactionTemplate(coreWorkflowVersion, rootPubKey, txType, SIGNATURES_TYPES.MULTISIG_SIG, nonce, fee, payload)

	multisigTransaction.Sig = aggregatedSignatureOfActive

	return multisigTransaction
}

func (web1337 *Web1337) BuildPartialSignatureWithTxData(originShard, txType, hexID string, sharedPayload []string, nonce uint, fee float32, payload map[string]interface{}) (string, error) {

	coreWorkflowVersion := web1337.Chains[web1337.CurrentChain].WorkflowVersion

	dataToSign := fmt.Sprintf("%s%d%s%s%s%d%d", web1337.CurrentChain, coreWorkflowVersion, originShard, txType, mapToJSON(payload), nonce, fee)

	partialSignature := tbls.GeneratePartialSignature(hexID, dataToSign, sharedPayload)

	return partialSignature, nil
}

func (sdk *Web1337) CreateThresholdTransaction(txType, tblsRootPubkey string, partialSignatures, idsOfSigners []string, nonce uint, fee float32, payload map[string]interface{}) TransactionTemplate {

	coreWorkflowVersion := sdk.Chains[sdk.CurrentChain].WorkflowVersion

	thresholdSigTransaction := sdk.GetTransactionTemplate(coreWorkflowVersion, tblsRootPubkey, txType, SIGNATURES_TYPES.TBLS_SIG, nonce, fee, payload)

	thresholdSigTransaction.Sig = tbls.BuildRootSignature(partialSignatures, idsOfSigners)

	return thresholdSigTransaction
}

func (sdk *Web1337) CreatePostQuantumTransaction(originShard, txType, pqcAlgorithm, yourAddress, yourPrivateKeyAsHex string, nonce uint, fee float32, payload map[string]interface{}) (TransactionTemplate, error) {

	coreWorkflowVersion := sdk.Chains[sdk.CurrentChain].WorkflowVersion

	var algoToAddToTx string

	if pqcAlgorithm == "bliss" {

		algoToAddToTx = SIGNATURES_TYPES.POST_QUANTUM_BLISS

	} else {

		algoToAddToTx = SIGNATURES_TYPES.POST_QUANTUM_DILITHIUM

	}

	transaction := sdk.GetTransactionTemplate(coreWorkflowVersion, yourAddress, txType, algoToAddToTx, nonce, fee, payload)

	if pqcAlgorithm == "bliss" {

		transaction.Sig = pqc.GenerateBlissSignature(yourPrivateKeyAsHex, fmt.Sprintf("%s%d%s%s%s%d%f", sdk.CurrentChain, coreWorkflowVersion, originShard, txType, payload, nonce, fee))

	} else {

		transaction.Sig = pqc.GenerateDilithiumSignature(yourPrivateKeyAsHex, fmt.Sprintf("%s%d%s%s%s%d%f", sdk.CurrentChain, coreWorkflowVersion, originShard, txType, payload, nonce, fee))

	}

	return transaction, nil
}
