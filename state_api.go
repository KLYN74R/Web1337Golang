package web1337

func (sdk *Web1337) GetDataFromState(shard string, cellID string) ([]byte, error) {

	return sdk.getRequest("/state/" + shard + "/" + cellID)

}

func (sdk *Web1337) GetPoolStats(poolId string) ([]byte, error) {

	return sdk.getRequest("/pool_stats/" + poolId)

}
