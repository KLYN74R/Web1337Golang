package tests

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"lukechampine.com/blake3"
)

func TestBlake3SimplePerformance(t *testing.T) {
	// Define the message to be hashed
	message := "43A+RQL4bDWVHQ/HwA2dWwdF5s1M9YDmgKRZ2EH/P/nA/6Cbb76n3eTh6DYUMLKaWS75uKxteAdaD4yxkDLvBw=="
	msgBytes := []byte(message)

	// Measure the time taken for hashing
	start := time.Now()
	hash := blake3.Sum256(msgBytes)
	duration := time.Since(start)

	// Print the hash in hexadecimal format
	fmt.Printf("Hash: %x\n", hash)

	// Print the time taken for the hashing operation
	fmt.Printf("Time taken to hash message: %v\n", duration)

	// Perform additional iterations to benchmark hashing performance
	var blake3Hash [32]byte
	iterations := 1000
	start = time.Now()
	for i := 0; i < iterations; i++ {
		blake3Hash = blake3.Sum256([]byte(fmt.Sprintf("%s%d", message, i)))
	}
	duration = time.Since(start)

	// Print the final hash and the time taken for multiple iterations
	fmt.Printf("Final hash after %d iterations: %s\n", iterations, hex.EncodeToString(blake3Hash[:]))
	fmt.Printf("Time taken for %d iterations: %v\n", iterations, duration)
}
