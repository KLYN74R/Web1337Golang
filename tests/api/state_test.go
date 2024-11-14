package tests

import (
	"fmt"
	"testing"

	web1337 "github.com/KLYN74R/Web1337Golang"
)

func TestGetDataFromState(t *testing.T) {
	myOptions := web1337.Options{
		NodeURL:         "http://localhost:7332",
		ChainID:         "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	shard := "shard_0"
	cellID := "0xb2ec32c9d7216163790ba3628a6a6b5a12db457c933b1f4627775b6dae468636233c6ad9931a8ef848a58353e60d33dd"

	data, err := sdkHandler.GetDataFromState(shard, cellID)
	if err != nil {
		t.Fatalf("Error fetching data from state: %v", err)
	} else {
		fmt.Println("Result: ", string(data))
	}
}

func TestGetPoolStats(t *testing.T) {
	myOptions := web1337.Options{
		NodeURL:         "http://localhost:7332",
		ChainID:         "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	poolID := "9GQ46rqY238rk2neSwgidap9ww5zbAN4dyqyC7j5ZnBK"

	poolStats, err := sdkHandler.GetPoolStats(poolID)
	if err != nil {
		t.Fatalf("Error fetching pool stats: %v", err)
	} else {
		fmt.Println("Result: ", string(poolStats))
	}
}

func TestGetAccountFromState(t *testing.T) {
	myOptions := web1337.Options{
		NodeURL:         "http://localhost:7332",
		ChainID:         "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	shard := "shard_0"
	accountID := "0x8f079049121d5e2ae885bdc6581df9fb68eab94a7aa3ae54bfe1d1ac35aceefbb202f656b0c1b56d64583630612a9970"

	accountData, err := sdkHandler.GetAccountFromState(shard, accountID)
	if err != nil {
		t.Fatalf("Error fetching account from state: %v", err)
	} else {
		fmt.Println("Result: ", string(accountData))
	}
}
