/*

Web1337 by KLY

For Golang devs

*/

package tests

import (
	"fmt"
	"testing"

	web1337 "github.com/KLYN74R/Web1337Golang"

	"github.com/KLYN74R/Web1337Golang/crypto_primitives"
)

func TestEd25519(t *testing.T) {

	if !(web1337.Ed25519Process()) {

		t.Error("Signature verification failed")

	}

}

func TestPQC(t *testing.T) {

	if !(web1337.DilithiumProcess() && web1337.BlissProcess()) {

		t.Error("Signature verification failed")

	}

}

func TestBls(t *testing.T) {

	if !(web1337.BlsProcess()) {

		t.Error("Signature verification failed")

	}

}

func TestTBLS(t *testing.T) {

	myIDs := crypto_primitives.GenerateTblsRandomIDs(6)

	fmt.Println("IDs => ", myIDs)

}
