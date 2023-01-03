package pkg2_test

import (
	"testing"

	"github.com/captain-examples/go-testing/internal/pkg2"
)

func TestFoo(t *testing.T) {
	if res := pkg2.Foo("arg"); res != "Foo: arg" {
		t.Errorf("%v != %v", res, "Foo: arg")
	}
}
