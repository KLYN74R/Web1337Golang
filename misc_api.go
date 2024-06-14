package web1337

func (sdk *Web1337) GetTargetNodeInfrastructureInfo() ([]byte, error) {

	return sdk.getRequest("/infrastructure_info")

}

func (sdk *Web1337) GetChainData() ([]byte, error) {

	return sdk.getRequest("/chain_info")

}

func (sdk *Web1337) GetKlyEvmMetadata() ([]byte, error) {

	return sdk.getRequest("/kly_evm_metadata")

}

func (sdk *Web1337) GetSynchronizationStatus() ([]byte, error) {

	return sdk.getRequest("/synchronization_stats")

}
