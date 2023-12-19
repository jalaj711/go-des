package godes

func xor64bit(a [8]byte, b [8]byte) (xored [8]byte) {
	for i := 0; i < 8; i++ {
		xored[i] = a[i] ^ b[i]
	}
	return xored
}

func Encrypt_CBC(plaintext []byte, key [8]byte, iv [8]byte) (ciphertext []byte) {
	plaintext = addPadding(plaintext)
	ciphertext = make([]byte, 0, len(plaintext))
	var block [8]byte
	for i := 0; i < len(plaintext); i += 8 {
		block = xor64bit(([8]byte)(plaintext[i:i+8]), iv)
		block = Encrypt64(block, key)
		iv = block
		ciphertext = append(ciphertext, block[:]...)
	}
	return ciphertext
}

func Decrypt_CBC(ciphertext []byte, key [8]byte, iv [8]byte) (plaintext []byte, err error) {
	plaintext = make([]byte, 0, len(ciphertext))
	var block [8]byte
	var ciphertextblock [8]byte
	for i := 0; i < len(ciphertext); i += 8 {
		ciphertextblock = ([8]byte)(ciphertext[i : i+8])
		block = xor64bit(iv, Decrypt64(ciphertextblock, key))
		iv = ciphertextblock
		plaintext = append(plaintext, block[:]...)
	}
	return removePadding(plaintext)
}
