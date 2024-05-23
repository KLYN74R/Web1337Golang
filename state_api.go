package web1337

func (sdk *Web1337) GetSyncState() ([]byte, error) {
	return sdk.getRequest("/sync_state")
}

func (sdk *Web1337) GetFromState(shard string, cellID string) ([]byte, error) {
	return sdk.getRequest("/state/" + shard + "/" + cellID)
}

func (sdk *Web1337) GetSymbioteInfo() ([]byte, error) {
	return sdk.getRequest("/symbiote_info")
}

func (sdk *Web1337) GetCurrentCheckpoint() ([]byte, error) {
	return sdk.getRequest("/quorum_thread_checkpoint")
}
