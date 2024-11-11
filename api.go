package web1337

import (
  "strconv"
)

/*
    Blocks
*/
func (sdk *Web1337) GetBlockByBlockID(blockID string) ([]byte, error) {
  return sdk.getRequest("/block/" + blockID)
}

func (sdk *Web1337) GetBlockBySID(shard string, indexInShard uint) ([]byte, error) {
    indexInShardStr := strconv.FormatUint(uint64(indexInShard), 10)

    return sdk.getRequest("/block_by_sid/" + shard + "/" + indexInShardStr)
}

func (sdk *Web1337) GetLatestNBlocksOnShard(shard string, startIndex uint, limit uint) ([]byte, error) {
    startIndexStr := strconv.FormatUint(uint64(startIndex), 10)
    limitStr := strconv.FormatUint(uint64(limit), 10)

    return sdk.getRequest("/latest_n_blocks/" + shard + "/" + startIndexStr + "/" + limitStr)
}

func (sdk *Web1337) GetTotalBlocksAndTxsStats() ([]byte, error) {
    return sdk.getRequest("/total_blocks_and_txs_stats")
}

/*
    Epochs
*/
func (sdk *Web1337) GetCurrentEpochOnThreads(threadID string) ([]byte, error) {
    return sdk.getRequest("/current_epoch/" + threadID)
}

func (sdk *Web1337) GetCurrentShardLeaders() ([]byte, error) {
    return sdk.getRequest("/current_shards_leaders")
}

func (sdk *Web1337) GetEpochByIndex(epochIndex uint) ([]byte, error) {
    epochIndexStr := strconv.FormatUint(uint64(epochIndex), 10)

    return sdk.getRequest("/epoch_by_index/" + epochIndexStr)
}

func (sdk *Web1337) GetTotalBlocksAndTxsByEpochIndex(epochIndex uint) ([]byte, error) {
    epochIndexStr := strconv.FormatUint(uint64(epochIndex), 10)

    return sdk.getRequest("/total_blocks_and_txs_stats_per_epoch/" + epochIndexStr)
}

func (sdk *Web1337) GetHistoricalStatsPerEpoch(startIndex uint, limit uint) ([]byte, error) {
    startIndexStr := strconv.FormatUint(uint64(startIndex), 10)
    limitStr := strconv.FormatUint(uint64(limit), 10)

    return sdk.getRequest("/historical_stats_per_epoch/" + startIndexStr + "/" + limitStr)
}

/*
    Transactions
*/
func (sdk *Web1337) GetTransactionReceipt(txID string) ([]byte, error) {
    return sdk.getRequest("/tx_receipt/" + txID)
}

func (sdk *Web1337) GetTransactionsList(shardID string, accountID string) ([]byte, error) {
    return sdk.getRequest("/txs_list/" + shardID + "/" + accountID)
}

/*
    State
*/
func (sdk *Web1337) GetDataFromState(shard string, cellID string) ([]byte, error) {
    return sdk.getRequest("/state/" + shard + "/" + cellID)
}

func (sdk *Web1337) GetPoolStats(poolID string) ([]byte, error) {
    return sdk.getRequest("/pool_stats/" + poolID)
}

func (sdk *Web1337) GetAccountFromState(shard string, accountID string) ([]byte, error) {
    return sdk.getRequest("/account/" + shard + "/" + accountID)
}

/*
    Consensus related
*/
func (sdk *Web1337) GetAggregatedFinalizationProof(blockID string) ([]byte, error) {
    return sdk.getRequest("/aggregated_finalization_proof/" + blockID)
}

func (sdk *Web1337) GetAggregatedEpochFinalizationProof(epochIndex uint, shard string) ([]byte, error) {
    epochIndexStr := strconv.FormatUint(uint64(epochIndex), 10)
    return sdk.getRequest("/aggregated_epoch_finalization_proof/" + epochIndexStr + "/" + shard)
}


/*
    Misc
*/
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