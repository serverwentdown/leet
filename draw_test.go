package main

import "testing"

func TestMixColors(t *testing.T) {
	mixed := mixColors(uint32(0), uint32(0xFFFF0000))
	if mixed != uint32(0xFFFF0000) {
		t.Errorf("mixColors(0, 0xFFFF0000) was incorrect, got: %#x, want: %#x", mixed, uint32(0xFFFF0000))
	}
	mixed = mixColors(uint32(0), uint32(0))
	if mixed != uint32(0) {
		t.Errorf("mixColors(0, 0) was incorrect, got: %#x, want: %#x", mixed, uint32(0))
	}
	mixed = mixColors(uint32(0x7FFF0000), uint32(0x7F00FF00))
	if mixed != uint32(0xbe7f3f00) {
		t.Errorf("mixColors(0x7FFF0000, 0x7F00FF00) was incorrect, got: %#x, want: %#x", mixed, uint32(0xbe7f3f00))
	}
}
