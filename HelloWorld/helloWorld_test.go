package helloworld_test

import (
	"testing"

	helloworld "github.com/SudeepthaG/golearn/HelloWorld"
)

func TestHello(t *testing.T) {
	if helloworld.Hello() != "Hello, World" { //Opening bracket should be present on same line as if

		// if the return string of helloworld.Hello() == "Hello, World" , trst will pass. otherwise it will fail.
		//Only return part matters for the test, every other print stmt inside Hello() is ignored here
		t.Fatal("Wrong welcome")
	}
}
