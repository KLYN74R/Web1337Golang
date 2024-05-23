package web1337

func (sdk *Web1337) GetAggregatedFinalizationProofForBlock(blockID string) ([]byte, error) {
	return sdk.getRequest("/aggregated_finalization_proof/" + blockID)
}

func (sdk *Web1337) GetGeneralInfoAboutKLYInfrastructure() ([]byte, error) {
	return sdk.getRequest("/my_info")
}
