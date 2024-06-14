package web1337

func (sdk *Web1337) GetAggregatedEpochFinalizationProof(epochIndex, shard string) ([]byte, error) {
	return sdk.getRequest("/aggregated_epoch_finalization_proof/" + epochIndex + "/" + shard)
}

func (sdk *Web1337) GetAggregatedFinalizationProofForBlock(blockId string) ([]byte, error) {
	return sdk.getRequest("/aggregated_finalization_proof/" + blockId)
}
