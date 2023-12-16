package godes

import (
	"testing"
)

func TestEncrypt64(t *testing.T) {
	test_key := [8]byte{0b00010011, 0b00110100, 0b01010111, 0b01111001, 0b10011011, 0b10111100, 0b11011111, 0b11110001}
	test_data := [8]byte{0b00000001, 0b00100011, 0b01000101, 0b01100111, 0b10001001, 0b10101011, 0b11001101, 0b11101111}
	expected := [8]byte{0b10000101, 0b11101000, 0b00010011, 0b01010100, 0b00001111, 0b00001010, 0b10110100, 0b00000101}
	encrypted := Encrypt64(test_data, test_key)
	if encrypted != expected {
		t.Fatalf(`Testcase: data=%b;key=%b, result: %b, expected: %b`, test_data, test_key, encrypted, expected)
	}
}

func TestDecrypt64(t *testing.T) {
	test_key := [8]byte{0b00010011, 0b00110100, 0b01010111, 0b01111001, 0b10011011, 0b10111100, 0b11011111, 0b11110001}
	test_data := [8]byte{0b10000101, 0b11101000, 0b00010011, 0b01010100, 0b00001111, 0b00001010, 0b10110100, 0b00000101}
	expected := [8]byte{0b00000001, 0b00100011, 0b01000101, 0b01100111, 0b10001001, 0b10101011, 0b11001101, 0b11101111}
	decrypted := Decrypt64(test_data, test_key)
	if decrypted != expected {
		t.Fatalf(`Testcase: data=%b;key=%b, result: %b, expected: %b`, test_data, test_key, decrypted, expected)
	}
}
