package tests

import (
    "fmt"
    "testing"
    web1337 "github.com/KLYN74R/Web1337Golang/pkg"
)

func TestGetInfrastructureInfo(t *testing.T) {
    myOptions := web1337.Options{
        NodeURL: "http://localhost:7332",
        SymbioteID: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
        WorkflowVersion: 1,
    }

    sdkHandler, _ := web1337.NewWeb1337(myOptions)

    infraInfo, err := sdkHandler.GetInfrastructureInfo()
    if err != nil {
        t.Fatalf("Error fetching infrastructure info: %v", err)
    } else {
        fmt.Println("Result: ", string(infraInfo))
    }
}

func TestGetChainInfo(t *testing.T) {
    myOptions := web1337.Options{
        NodeURL: "http://localhost:7332",
        SymbioteID: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
        WorkflowVersion: 1,
    }

    sdkHandler, _ := web1337.NewWeb1337(myOptions)

    chainInfo, err := sdkHandler.GetChainInfo()
    if err != nil {
        t.Fatalf("Error fetching chain info: %v", err)
    } else {
        fmt.Println("Result: ", string(chainInfo))
    }
}

func TestGetKlyEVMMetadata(t *testing.T) {
    myOptions := web1337.Options{
        NodeURL: "http://localhost:7332",
        SymbioteID: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
        WorkflowVersion: 1,
    }

    sdkHandler, _ := web1337.NewWeb1337(myOptions)

    evmMetadata, err := sdkHandler.GetKlyEVMMetadata()
    if err != nil {
        t.Fatalf("Error fetching Kly EVM metadata: %v", err)
    } else {
        fmt.Println("Result: ", string(evmMetadata))
    }
}

func TestGetSynchronizationStats(t *testing.T) {
    myOptions := web1337.Options{
        NodeURL: "http://localhost:7332",
        SymbioteID: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
        WorkflowVersion: 1,
    }

    sdkHandler, _ := web1337.NewWeb1337(myOptions)

    syncStats, err := sdkHandler.GetSynchronizationStats()
    if err != nil {
        t.Fatalf("Error fetching synchronization stats: %v", err)
    } else {
        fmt.Println("Result: ", string(syncStats))
    }
}

// func TestGetCheckpointByEpochIndex(t *testing.T) {
//     myOptions := web1337.Options{
//         NodeURL: "http://localhost:7332",
//         SymbioteID: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
//         WorkflowVersion: 1,
//     }
//
//     sdkHandler, _ := web1337.NewWeb1337(myOptions)
//
//     epochIndex := uint(0)
//     checkpointData, err := sdkHandler.GetCheckpointByEpochIndex(epochIndex)
//     if err != nil {
//         t.Fatalf("Error fetching checkpoint by epoch index: %v", err)
//     } else {
//         fmt.Println("Result: ", string(checkpointData))
//     }
// }
