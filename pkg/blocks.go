package web1337

import (
    "strconv"
)

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
