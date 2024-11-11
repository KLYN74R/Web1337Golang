package web1337

func (sdk *Web1337) GetDataFromState(shard string, cellID string) ([]byte, error) {
    return sdk.getRequest("/state/" + shard + "/" + cellID)
}

func (sdk *Web1337) GetPoolStats(poolID string) ([]byte, error) {
    return sdk.getRequest("/pool_stats/" + poolID)
}

func (sdk *Web1337) GetAccountFromState(shard string, accountID string) ([]byte, error) {
    return sdk.getRequest("/account/" + shard + "/" + accountID)
}

