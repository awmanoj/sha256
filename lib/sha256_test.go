package lib

import (
	"testing"
)

func TestNewSHA256Module(t *testing.T) {
	module := NewSHA256Module()
	if module == nil {
		t.Errorf("err TestNewSHA256Module, module is nil")
	}
}

func TestEncode(t *testing.T) {
	module := NewSHA256Module()
	actual := module.SHA256("What's the answer to the universe?", "42")
	expected := "46f9d41a494d1b8009939f4e3199be05d88c59fb378c9e64b66d980f0bd20cf4"
	if actual != expected {
		t.Errorf("err TestEncode, mismatch in the SHA256 encoded string")
	}
}
