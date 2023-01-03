package pkg1_test

import (
	"testing"

	"github.com/captain-examples/go-testing/internal/pkg1"
)

func TestFoo(t *testing.T) {
	if res := pkg1.Foo("arg"); res != "Foo: arg" {
		t.Errorf("%v != %v", res, "Foo: arg")
	}
}
