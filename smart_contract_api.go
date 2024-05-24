package web1337

func (sdk *Web1337) GetContractMetadata(contractId string) ([]byte, error) {
	return sdk.getRequest("/account/" + contractId)
}

func (sdk *Web1337) GetContractStorage(contractId, storageName string) ([]byte, error) {
	return sdk.getRequest("/account/" + contractId + "_STORAGE_" + storageName)
}

func (sdk *Web1337) CallContract(contractId, method string, params map[string]interface{}, injects []string) {
	// return sdk.getRequest("/symbiote_info")
}

func (sdk *Web1337) deployContractToWvm(bytecode string) {
	// return sdk.getRequest("/quorum_thread_checkpoint")
}
