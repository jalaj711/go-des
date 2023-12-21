package godes

import "errors"

func TripleDES128_Encrypt64(plaintext [8]byte, key [16]byte) (ciphertext [8]byte) {
	K1 := [8]byte(key[:8])
	K2 := [8]byte(key[8:])
	return Encrypt64(Decrypt64(Encrypt64(plaintext, K1), K2), K1)
}

func TripleDES128_Decrypt64(plaintext [8]byte, key [16]byte) (ciphertext [8]byte) {
	K1 := [8]byte(key[:8])
	K2 := [8]byte(key[8:])
	return Decrypt64(Encrypt64(Decrypt64(plaintext, K1), K2), K1)
}

func TripleDES192_Encrypt64(plaintext [8]byte, key [24]byte) (ciphertext [8]byte) {
	K1 := [8]byte(key[:8])
	K2 := [8]byte(key[8:16])
	K3 := [8]byte(key[16:])
	return Encrypt64(Decrypt64(Encrypt64(plaintext, K1), K2), K3)
}

func TripleDES192_Decrypt64(plaintext [8]byte, key [24]byte) (ciphertext [8]byte) {
	K1 := [8]byte(key[:8])
	K2 := [8]byte(key[8:16])
	K3 := [8]byte(key[16:])
	return Decrypt64(Encrypt64(Decrypt64(plaintext, K1), K2), K3)
}

func TripleDES128_Encrypt(data []byte, key [16]byte) (encrypted []byte) {
	data = addPadding(data)
	encrypted = make([]byte, 0, len(data))
	var block [8]byte
	for i := 0; i < len(data); i += 8 {
		block = (TripleDES128_Encrypt64(([8]byte)(data[i:i+8]), key))
		encrypted = append(encrypted, block[:]...)
	}
	return encrypted
}

func TripleDES128_Decrypt(data []byte, key [16]byte) (decrypted []byte, err error) {
	if len(data)%8 != 0 {
		return decrypted, errors.New("ciphertext must be a multiple of 64-bits")
	}
	decrypted = make([]byte, 0, len(data))
	var block [8]byte
	for i := 0; i < len(data); i += 8 {
		block = (TripleDES128_Decrypt64(([8]byte)(data[i:i+8]), key))
		decrypted = append(decrypted, block[:]...)
	}
	return removePadding(decrypted)
}

func TripleDES192_Encrypt(data []byte, key [24]byte) (encrypted []byte) {
	data = addPadding(data)
	encrypted = make([]byte, 0, len(data))
	var block [8]byte
	for i := 0; i < len(data); i += 8 {
		block = (TripleDES192_Encrypt64(([8]byte)(data[i:i+8]), key))
		encrypted = append(encrypted, block[:]...)
	}
	return encrypted
}

func TripleDES192_Decrypt(data []byte, key [24]byte) (decrypted []byte, err error) {
	if len(data)%8 != 0 {
		return decrypted, errors.New("ciphertext must be a multiple of 64-bits")
	}
	decrypted = make([]byte, 0, len(data))
	var block [8]byte
	for i := 0; i < len(data); i += 8 {
		block = (TripleDES192_Decrypt64(([8]byte)(data[i:i+8]), key))
		decrypted = append(decrypted, block[:]...)
	}
	return removePadding(decrypted)
}
