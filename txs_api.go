package web1337

func (sdk *Web1337) GetTransactionReceiptById(txID string) ([]byte, error) {
	return sdk.getRequest("/tx_receipt/" + txID)
}

func (sdk *Web1337) SendTransaction(transaction TransactionTemplate) ([]byte, error) {
	return sdk.postRequest("/transaction", transaction)
}
