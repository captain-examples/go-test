package pkg1_test

import (
	"testing"

	"github.com/captain-examples/go-testing/internal/pkg1"
)

func TestBar(t *testing.T) {
	if res := pkg1.Bar("arg"); res != "Bar: arg" {
		t.Errorf("%v != %v", res, "Bar: arg")
	}
}
