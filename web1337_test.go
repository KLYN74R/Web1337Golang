/*

Web1337 by KLY

For Golang devs

*/

package web1337

import (
	"testing"
)

func TestHello(t *testing.T) {

	if Hello() != "Hello from Web1337" {

		t.Errorf("Hello() func didn't return greeting")

	}

}
