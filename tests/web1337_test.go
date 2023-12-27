/*

Web1337 by KLY

For Golang devs

*/

package tests

import (
	"testing"

	web1337 "github.com/KLYN74R/Web1337Golang"
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

func TestTed25519(t *testing.T) {

	if !(web1337.Ted25519Process()) {

		t.Error("Test failed")

	}

}
