package godes

import (
	"testing"
)

func TestExpansionFuncton(t *testing.T) {
	test := [4]byte{0b11110000, 0b10101010, 0b11110000, 0b10101010}
	expected := [6]byte{0b01111010, 0b00010101, 0b01010101, 0b01111010, 0b00010101, 0b01010101}
	key := expand(test)
	// t.Logf(`Testcase: %b, result: %d`, test, key)
	if key != expected {
		t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
	}
}
