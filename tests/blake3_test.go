package tests

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"

	"lukechampine.com/blake3"
)

func TestBlake3SimplePerformance(t *testing.T) {

	// msg := []byte("Hello")

	var blake3Hash [32]byte

	for i := 0; i < 1000000; i++ {

		msg := []byte("Hello" + strconv.Itoa(i))

		blake3Hash = blake3.Sum256(msg)

	}

	fmt.Println(hex.EncodeToString(blake3Hash[:]))

	fmt.Printf("%x", blake3Hash)

}
