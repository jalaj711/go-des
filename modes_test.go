package godes

import "testing"

func TestEncrypt_CBC(t *testing.T) {
	test_key := [8]byte{0xda, 0x39, 0xa3, 0xee, 0x5e, 0x6b, 0x4b, 0x0d}
	test_data := []byte("12345678")
	test_iv := [8]byte{0xff, 0x40, 0x48, 0xa9, 0xca, 0xe7, 0x6b, 0xbe}
	expected := []byte{0x1c, 0x33, 0xac, 0x21, 0x57, 0x56, 0xb2, 0xcd, 0x2c, 0x91, 0x92, 0x6f, 0xd9, 0x0b, 0x77, 0x20}
	encrypted := Encrypt_CBC(test_data, test_key, test_iv)
	if len(encrypted) != len(expected) {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, test_data, test_key, encrypted, expected)
	}
	for i := range expected {
		if encrypted[i] != expected[i] {
			t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x, diff=%d`, test_data, test_key, encrypted, expected, i)
		}
	}
}

func TestDecrypt_CBC(t *testing.T) {
	test_key := [8]byte{0xda, 0x39, 0xa3, 0xee, 0x5e, 0x6b, 0x4b, 0x0d}
	test_data := []byte{0x1c, 0x33, 0xac, 0x21, 0x57, 0x56, 0xb2, 0xcd, 0x2c, 0x91, 0x92, 0x6f, 0xd9, 0x0b, 0x77, 0x20}
	test_iv := [8]byte{0xff, 0x40, 0x48, 0xa9, 0xca, 0xe7, 0x6b, 0xbe}
	expected := []byte("12345678")
	decrypted, err := Decrypt_CBC(test_data, test_key, test_iv)
	if err != nil {
		t.Fatalf(`Testcase: data=%x;key=%x, result: error: %s, expected: %x`, test_data, test_key, err.Error(), expected)
	}
	if len(decrypted) != len(expected) {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, test_data, test_key, decrypted, expected)
	}
	for i := range expected {
		if decrypted[i] != expected[i] {
			t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x, diff=%d`, test_data, test_key, decrypted, expected, i)
		}
	}
}

func TestEncrypt_CFB8(t *testing.T) {
	test_key := [8]byte{0xda, 0x39, 0xa3, 0xee, 0x5e, 0x6b, 0x4b, 0x0d}
	test_data := []byte("12345678")
	test_iv := [8]byte{0xff, 0x40, 0x48, 0xa9, 0xca, 0xe7, 0x6b, 0xbe}

	expected := []byte{0x2c, 0x2b, 0xfb, 0xd3, 0xe0, 0x0b, 0xc5, 0xb4, 0xec, 0xfb, 0x8b, 0xd9, 0xb3, 0xac, 0xa9, 0xfe}
	encrypted, _ := Encrypt_CFB8(test_data, test_key, test_iv)
	if len(encrypted) != len(expected) {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, test_data, test_key, encrypted, expected)
	}
	for i := range expected {
		if encrypted[i] != expected[i] {
			t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x, diff=%d`, test_data, test_key, encrypted, expected, i)
		}
	}
}

func TestDecrypt_CFB8(t *testing.T) {
	test_key := [8]byte{0xda, 0x39, 0xa3, 0xee, 0x5e, 0x6b, 0x4b, 0x0d}
	test_data := []byte{0x2c, 0x2b, 0xfb, 0xd3, 0xe0, 0x0b, 0xc5, 0xb4, 0xec, 0xfb, 0x8b, 0xd9, 0xb3, 0xac, 0xa9, 0xfe}
	test_iv := [8]byte{0xff, 0x40, 0x48, 0xa9, 0xca, 0xe7, 0x6b, 0xbe}

	expected := []byte("12345678")
	encrypted, _ := Decrypt_CFB8(test_data, test_key, test_iv)
	if len(encrypted) != len(expected) {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, test_data, test_key, encrypted, expected)
	}
	for i := range expected {
		if encrypted[i] != expected[i] {
			t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x, diff=%d`, test_data, test_key, encrypted, expected, i)
		}
	}
}
