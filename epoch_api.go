package web1337

func (sdk *Web1337) GetCurrentEpochOnThread(threadID string) ([]byte, error) {

	return sdk.getRequest("/current_epoch/" + threadID)

}

func (sdk *Web1337) GetCurrentLeadersOnShards() ([]byte, error) {

	return sdk.getRequest("/current_shards_leaders")

}

func (sdk *Web1337) GetEpochDataByEpochIndex(epochIndex string) ([]byte, error) {

	return sdk.getRequest("/epoch_by_index/" + epochIndex)

}
