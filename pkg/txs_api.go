package web1337

func (sdk *Web1337) GetTransactionReceipt(txID string) ([]byte, error) {
    return sdk.getRequest("/tx_receipt/" + txID)
}

func (sdk *Web1337) GetTransactionsList(shardID string, accountID string) ([]byte, error) {
    return sdk.getRequest("/txs_list/" + shardID + "/" + accountID)
}
