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

func TestEncrypt(t *testing.T) {
	test_key := []byte("64bitKey")
	test_data := []byte("1234567812345678")
	expected := []byte{0xa0, 0xbc, 0x3d, 0x48, 0xfa, 0x79, 0xda, 0xfb, 0xa0, 0xbc, 0x3d, 0x48, 0xfa, 0x79, 0xda, 0xfb, 0xb5, 0xd1, 0xba, 0xd5, 0x84, 0xf1, 0x30, 0x7c}
	encrypted := Encrypt(test_data, [8]byte(test_key))
	if len(encrypted) != len(expected) {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, test_data, test_key, encrypted, expected)
	}
	for i := range expected {
		if encrypted[i] != expected[i] {
			t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x, diff=%d`, test_data, test_key, encrypted, expected, i)
		}
	}
}

func TestDecrypt(t *testing.T) {
	test_key := []byte("64bitKey")
	test_data := []byte{0xa0, 0xbc, 0x3d, 0x48, 0xfa, 0x79, 0xda, 0xfb, 0xa0, 0xbc, 0x3d, 0x48, 0xfa, 0x79, 0xda, 0xfb, 0xb5, 0xd1, 0xba, 0xd5, 0x84, 0xf1, 0x30, 0x7c}
	expected := []byte("1234567812345678")
	decrypted, _ := Decrypt(test_data, [8]byte(test_key))
	if len(decrypted) != len(expected) {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, test_data, test_key, decrypted, expected)
	}
	for i := range expected {
		if decrypted[i] != expected[i] {
			t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x, diff=%d`, test_data, test_key, decrypted, expected, i)
		}
	}
}

func TestDecryptInvalidPadding(t *testing.T) {
	test_key := []byte("64bitKey")
	test_data := []byte{0x53, 0x62, 0x38, 0x3a, 0x75, 0x2c, 0xf4, 0xdf}
	_, err := Decrypt(test_data, [8]byte(test_key))
	if err == nil {
		t.Fatalf(`Testcase: data=%x;key=%x, expected error`, test_data, test_key)
	}
}

func TestDecryptInvalidLength(t *testing.T) {
	test_key := []byte("64bitKey")
	test_data := []byte{0x53, 0x62, 0x38, 0x3a, 0x75, 0x2c, 0xf4}
	_, err := Decrypt(test_data, [8]byte(test_key))
	if err == nil {
		t.Fatalf(`Testcase: data=%x;key=%x, expected error`, test_data, test_key)
	}
}
