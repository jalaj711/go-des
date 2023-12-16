package godes

import (
	"testing"
)

func TestGetReducedKey(t *testing.T) {
	test := [8]byte{0b00010011, 0b00110100, 0b01010111, 0b01111001, 0b10011011, 0b10111100, 0b11011111, 0b11110001}
	expected := [7]byte{0b11110000, 0b11001100, 0b10101010, 0b11110101, 0b01010110, 0b01100111, 0b10001111}
	key := getReducedKey(test)
	// t.Logf(`Testcase: %b, result: %d`, test, key)
	if key != expected {
		t.Fatalf(`Testcase: %b, result: %b`, test, key)
	}
}
