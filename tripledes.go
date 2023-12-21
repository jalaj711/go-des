package godes

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