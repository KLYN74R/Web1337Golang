package web1337

import (
    "strconv"
)

func (sdk *Web1337) GetInfrastructureInfo() ([]byte, error) {
    return sdk.getRequest("/infrastructure_info")
}

func (sdk *Web1337) GetChainInfo() ([]byte, error) {
    return sdk.getRequest("/chain_info")
}

func (sdk *Web1337) GetKlyEVMMetadata() ([]byte, error) {
    return sdk.getRequest("/kly_evm_metadata")
}

func (sdk *Web1337) GetSynchronizationStats() ([]byte, error) {
    return sdk.getRequest("/synchronization_stats")
}

func (sdk *Web1337) GetCheckpointByEpochIndex(epochIndex uint) ([]byte, error) {
    epochIndexStr := strconv.FormatUint(uint64(epochIndex), 10)
    return sdk.getRequest("/checkpoints/" + epochIndexStr)
}
