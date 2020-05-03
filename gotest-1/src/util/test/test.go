package test

import (
	"testing"
)

func Assert(t *testing.T, b bool, have interface{}, want interface{}) {
	if !b {
		t.Errorf("Assertion failed! have: %v, want %v", have, want)
	}
}
