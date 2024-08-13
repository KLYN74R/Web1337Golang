package web1337

import (
	"fmt"

	ed25519 "github.com/KLYN74R/Web1337Golang/crypto_primitives/ed25519"
	"github.com/KLYN74R/Web1337Golang/crypto_primitives/pqc"
	tbls "github.com/KLYN74R/Web1337Golang/crypto_primitives/tbls"

	SIGNATURES_TYPES "github.com/KLYN74R/Web1337Golang/signatures_types"
	TXS_TYPES "github.com/KLYN74R/Web1337Golang/txs_types"
)

type TransactionTemplate struct {
	V       uint        `json:"v"`
	Creator string      `json:"creator"`
	Type    string      `json:"type"`
	Nonce   uint        `json:"nonce"`
	Fee     float32     `json:"fee"`
	Payload interface{} `json:"payload"` // might be various - depends on transaction type
	Sig     string      `json:"sig"`
}

func (web1337 *Web1337) GetTransactionTemplate(workflowVersion uint, creator, txType string, nonce uint, fee float32, payload map[string]interface{}) TransactionTemplate {

	return TransactionTemplate{
		V:       workflowVersion,
		Creator: creator,
		Type:    txType,
		Nonce:   nonce,
		Fee:     fee,
		Payload: payload,
		Sig:     "",
	}
}

func (web1337 *Web1337) CreateDefaultTransaction(originShard, yourAddress, yourPrivateKeyAsHex string, nonce uint, fee float32, recipient string, amountInKLY float32, rev_t *int) (TransactionTemplate, error) {

	coreWorkflowVersion := web1337.Symbiotes[web1337.CurrentSymbiote].WorkflowVersion

	payload := map[string]interface{}{
		"sigType": SIGNATURES_TYPES.DEFAULT_SIG,
		"to":      recipient,
		"amount":  amountInKLY,
	}

	// In case we send from Ed25519 to BLS
	if rev_t != nil {
		payload["rev_t"] = *rev_t
	}

	txTemplate := web1337.GetTransactionTemplate(coreWorkflowVersion, yourAddress, TXS_TYPES.TX, nonce, fee, payload)

	dataToSign := web1337.CurrentSymbiote + string(coreWorkflowVersion) + originShard + TXS_TYPES.TX + mapToJSON(payload) + string(nonce) + fmt.Sprintf("%f", fee)

	txTemplate.Sig = ed25519.GenerateSignature(yourPrivateKeyAsHex, dataToSign)

	return txTemplate, nil
}

func (web1337 *Web1337) CreateMultisigTransaction(rootPubKey, aggregatedPubOfActive, aggregatedSignatureOfActive string, afkSigners []string, nonce uint, fee float32, recipient string, amountInKLY float32, rev_t *int) TransactionTemplate {

	coreWorkflowVersion := web1337.Symbiotes[web1337.CurrentSymbiote].WorkflowVersion

	payload := map[string]interface{}{
		"sigType": SIGNATURES_TYPES.MULTISIG_SIG,
		"active":  aggregatedPubOfActive,
		"afk":     afkSigners,
		"to":      recipient,
		"amount":  amountInKLY,
	}

	if rev_t != nil {
		payload["rev_t"] = *rev_t
	}

	multisigTransaction := web1337.GetTransactionTemplate(coreWorkflowVersion, rootPubKey, TXS_TYPES.TX, nonce, fee, payload)
	multisigTransaction.Sig = aggregatedSignatureOfActive

	return multisigTransaction
}

// BuildPartialSignatureWithTxData builds a partial signature with transaction data
func (web1337 *Web1337) BuildPartialSignatureWithTxData(hexID string, sharedPayload []string, originShard string, nonce, fee int, recipient string, amountInKLY int, rev_t *int) (string, error) {

	coreWorkflowVersion := web1337.Symbiotes[web1337.CurrentSymbiote].WorkflowVersion

	payloadForTblsTransaction := map[string]interface{}{
		"to":      recipient,
		"amount":  amountInKLY,
		"sigType": SIGNATURES_TYPES.TBLS_SIG,
	}

	if rev_t != nil {
		payloadForTblsTransaction["rev_t"] = *rev_t
	}

	dataToSign := fmt.Sprintf("%s%d%s%s%s%d%d", web1337.CurrentSymbiote, coreWorkflowVersion, originShard, TXS_TYPES.TX, mapToJSON(payloadForTblsTransaction), nonce, fee)
	partialSignature := tbls.GeneratePartialSignature(hexID, dataToSign, sharedPayload)

	return partialSignature, nil
}

// CreateThresholdTransaction creates a threshold transaction
func (sdk *Web1337) CreateThresholdTransaction(tblsRootPubkey string, partialSignatures, idsOfSigners []string, nonce uint, recipient string, amountInKLY, fee float32, rev_t *int) TransactionTemplate {

	coreWorkflowVersion := sdk.Symbiotes[sdk.CurrentSymbiote].WorkflowVersion

	tblsPayload := map[string]interface{}{
		"to":      recipient,
		"amount":  amountInKLY,
		"sigType": SIGNATURES_TYPES.TBLS_SIG,
	}

	if rev_t != nil {
		tblsPayload["rev_t"] = *rev_t
	}

	thresholdSigTransaction := sdk.GetTransactionTemplate(coreWorkflowVersion, tblsRootPubkey, TXS_TYPES.TX, nonce, fee, tblsPayload)
	thresholdSigTransaction.Sig = tbls.BuildRootSignature(partialSignatures, idsOfSigners)

	return thresholdSigTransaction
}

// CreatePostQuantumTransaction creates a post-quantum transaction
func (sdk *Web1337) CreatePostQuantumTransaction(originShard, sigType, yourAddress, yourPrivateKeyAsHex string, nonce uint, recipient string, amountInKLY, fee float32, rev_t *int) (TransactionTemplate, error) {

	coreWorkflowVersion := sdk.Symbiotes[sdk.CurrentSymbiote].WorkflowVersion

	payload := map[string]interface{}{
		"sigType": sigType,
		"to":      recipient,
		"amount":  amountInKLY,
	}

	if rev_t != nil {
		payload["rev_t"] = *rev_t
	}

	transaction := sdk.GetTransactionTemplate(coreWorkflowVersion, yourAddress, TXS_TYPES.TX, nonce, fee, payload)

	if sigType == SIGNATURES_TYPES.POST_QUANTUM_BLISS {

		transaction.Sig = pqc.GenerateBlissSignature(yourPrivateKeyAsHex, fmt.Sprintf("%s%d%s%s%s%d%f", sdk.CurrentSymbiote, coreWorkflowVersion, originShard, TXS_TYPES.TX, payload, nonce, fee))

	} else {

		transaction.Sig = pqc.GenerateDilithiumSignature(yourPrivateKeyAsHex, fmt.Sprintf("%s%d%s%s%s%d%f", sdk.CurrentSymbiote, coreWorkflowVersion, originShard, TXS_TYPES.TX, payload, nonce, fee))

	}

	return transaction, nil
}
