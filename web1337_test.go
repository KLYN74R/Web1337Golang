/*

Web1337 by KLY

For Golang devs

*/

package web1337

import (
	"testing"
)

func TestEd25519(t *testing.T) {

	if !(Ed25519Process()) {

		t.Error("Signature verification failed")

	}

}

func TestPQC(t *testing.T) {

	if !(DilithiumProcess() && BlissProcess()) {

		t.Error("Signature verification failed")

	}

}

func TestTed25519(t *testing.T) {

	if !(Ted25519Process()) {

		t.Error("Test failed")

	}

}
