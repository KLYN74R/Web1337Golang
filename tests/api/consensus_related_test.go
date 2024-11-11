package tests

import (
    "fmt"
    "testing"
    web1337 "github.com/KLYN74R/Web1337Golang"
)

func TestGetAggregatedFinalizationProof(t *testing.T) {
    myOptions := web1337.Options{
        NodeURL: "http://localhost:7332",
        SymbioteID: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
        WorkflowVersion: 1,
    }

    sdkHandler, _ := web1337.NewWeb1337(myOptions)

    blockID := "0:9GQ46rqY238rk2neSwgidap9ww5zbAN4dyqyC7j5ZnBK:700"

    finalizationProof, err := sdkHandler.GetAggregatedFinalizationProof(blockID)
    if err != nil {
        t.Fatalf("Error fetching aggregated finalization proof: %v", err)
    } else {
        fmt.Println("Result: ", string(finalizationProof))
    }
}

func TestGetAggregatedEpochFinalizationProof(t *testing.T) {
    myOptions := web1337.Options{
        NodeURL: "http://localhost:7332",
        SymbioteID: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
        WorkflowVersion: 1,
    }

    sdkHandler, _ := web1337.NewWeb1337(myOptions)

    epochIndex := uint(0)
    shard := "shard_0"

    epochFinalizationProof, err := sdkHandler.GetAggregatedEpochFinalizationProof(epochIndex, shard)
    if err != nil {
        t.Fatalf("Error fetching aggregated epoch finalization proof: %v", err)
    } else {
        fmt.Println("Result: ", string(epochFinalizationProof))
    }
}
