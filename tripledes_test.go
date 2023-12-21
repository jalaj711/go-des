package godes

import "testing"

func TestTripleDES128_Encrypt64(t *testing.T) {
	test_key := [16]byte([]byte("1234567812345678"))
	test_data := [8]byte([]byte("12345678"))
	expected := [8]byte{0x96, 0xd0, 0x02, 0x88, 0x78, 0xd5, 0x8c, 0x89}
	encrypted := TripleDES128_Encrypt64(test_data, test_key)
	if encrypted != expected {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, test_data, test_key, encrypted, expected)
	}
}

func TestTripleDES128_Decrypt64(t *testing.T) {
	test_key := [16]byte([]byte("1234567812345678"))
	test_data := [8]byte{0x96, 0xd0, 0x02, 0x88, 0x78, 0xd5, 0x8c, 0x89}
	expected := [8]byte([]byte("12345678"))
	decrypted := TripleDES128_Decrypt64(test_data, test_key)
	if decrypted != expected {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, test_data, test_key, decrypted, expected)
	}
}
