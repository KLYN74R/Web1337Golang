package web1337

import (
    "strconv"
)

func (sdk *Web1337) GetAggregatedFinalizationProof(blockID string) ([]byte, error) {
    return sdk.getRequest("/aggregated_finalization_proof/" + blockID)
}

func (sdk *Web1337) GetAggregatedEpochFinalizationProof(epochIndex uint, shard string) ([]byte, error) {
    epochIndexStr := strconv.FormatUint(uint64(epochIndex), 10)
    return sdk.getRequest("/aggregated_epoch_finalization_proof/" + epochIndexStr + "/" + shard)
}
