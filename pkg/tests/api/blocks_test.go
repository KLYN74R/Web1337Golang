package tests

import (
	"fmt"
	"testing"
	web1337 "github.com/KLYN74R/Web1337Golang/pkg"
)

func TestGetBlockByBlockID(t *testing.T) {
    myOptions := web1337.Options{
        NodeURL: "http://localhost:7332",
		SymbioteID: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	blockID := "0:9GQ46rqY238rk2neSwgidap9ww5zbAN4dyqyC7j5ZnBK:15"

	blockData, err := sdkHandler.GetBlockByBlockID(blockID)
    if err != nil {
        t.Fatalf("Error fetching block by ID: %v", err)
    }

    fmt.Println("Result: ", string(blockData))
}

func TestGetBlockBySID(t *testing.T) {
    myOptions := web1337.Options{
        NodeURL: "http://localhost:7332",
		SymbioteID: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	shard := "shard_0"
	indexInShard := uint(1000)

	blockData, err := sdkHandler.GetBlockBySID(shard, indexInShard)

    if err != nil {
        t.Fatalf("Error fetching block by SID: %v", err)
    }

    fmt.Println("Result: ", string(blockData))
}

func TestGetLatestNBlocksOnShard(t *testing.T) {
    myOptions := web1337.Options{
        NodeURL: "http://localhost:7332",
		SymbioteID: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

    shard := "shard_0"
    startIndex := uint(100)
    limit := uint(10)

	blockData, err := sdkHandler.GetLatestNBlocksOnShard(shard, startIndex, limit)

    if err != nil {
        t.Fatalf("Error fetching latest N blocks on shard: %v", err)
    }

    fmt.Println("Result: ", string(blockData))
}

func TestGetTotalBlocksAndTxsStats(t *testing.T) {
    myOptions := web1337.Options{
        NodeURL: "http://localhost:7332",
		SymbioteID: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	blockAndTxInfo, err := sdkHandler.GetTotalBlocksAndTxsStats()

	if err != nil {
        t.Fatalf("Error fetching total blocks and transactions stats: %v", err)
    }

    fmt.Println("Result: ", string(blockAndTxInfo))
}