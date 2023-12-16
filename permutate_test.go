package godes

import (
	"testing"
)

func TestInitialPermutation(t *testing.T) {
	test := [8]byte{0b00000001, 0b00100011, 0b01000101, 0b01100111, 0b10001001, 0b10101011, 0b11001101, 0b11101111}
	expected := [8]byte{0b11001100, 0b00000000, 0b11001100, 0b11111111, 0b11110000, 0b10101010, 0b11110000, 0b10101010}
	key := initial_permutation(test)
	// t.Logf(`Testcase: %b, result: %d`, test, key)
	if key != expected {
		t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
	}
}

func TestInverseInitialPermutation(t *testing.T) {
	test := [8]byte{0b00001010, 0b01001100, 0b11011001, 0b10010101, 0b01000011, 0b01000010, 0b00110010, 0b00110100}
	expected := [8]byte{0b10000101, 0b11101000, 0b00010011, 0b01010100, 0b00001111, 0b00001010, 0b10110100, 0b00000101}
	key := inverse_initial_permutation(test)
	// t.Logf(`Testcase: %b, result: %d`, test, key)
	if key != expected {
		t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
	}
}
