package web1337

func (sdk *Web1337) GetBlockByBlockID(blockID string) ([]byte, error) {
	return sdk.getRequest("/block/" + blockID)
}

func (sdk *Web1337) GetBlockBySID(shard string, sid string) ([]byte, error) {
	return sdk.getRequest("/block_by_sid/" + shard + "/" + sid)
}
