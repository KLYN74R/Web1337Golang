package web1337

import (
    "strconv"
)

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
