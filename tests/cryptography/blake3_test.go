package tests

import (
	"fmt"
	"testing"

	"lukechampine.com/blake3"
)

func TestBlake3SimplePerformance(t *testing.T) {

	// msg := []byte("Hello")

	fmt.Printf("%x", blake3.Sum256([]byte("43A+RQL4bDWVHQ/HwA2dWwdF5s1M9YDmgKRZ2EH/P/nA/6Cbb76n3eTh6DYUMLKaWS75uKxteAdaD4yxkDLvBw==")))

	// var blake3Hash [32]byte

	// for i := 0; i < 1; i++ {

	// 	msg := []byte("Hello" + strconv.Itoa(i))

	// 	blake3Hash = blake3.Sum256(msg)

	// }

	// fmt.Println(hex.EncodeToString(blake3Hash[:]))

	// fmt.Printf("%x", blake3Hash)

}
