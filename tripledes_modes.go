package godes

import "errors"

func TripleDES128_Encrypt_CBC(plaintext []byte, key [16]byte, iv [8]byte) (ciphertext []byte) {
	plaintext = addPadding(plaintext)
	ciphertext = make([]byte, 0, len(plaintext))
	var block [8]byte
	for i := 0; i < len(plaintext); i += 8 {
		block = xor64bit(([8]byte)(plaintext[i:i+8]), iv)
		block = TripleDES128_Encrypt64(block, key)
		iv = block
		ciphertext = append(ciphertext, block[:]...)
	}
	return ciphertext
}

func TripleDES128_Decrypt_CBC(ciphertext []byte, key [16]byte, iv [8]byte) (plaintext []byte, err error) {
	plaintext = make([]byte, 0, len(ciphertext))
	var block [8]byte
	var ciphertextblock [8]byte
	for i := 0; i < len(ciphertext); i += 8 {
		ciphertextblock = ([8]byte)(ciphertext[i : i+8])
		block = xor64bit(iv, TripleDES128_Decrypt64(ciphertextblock, key))
		iv = ciphertextblock
		plaintext = append(plaintext, block[:]...)
	}
	return removePadding(plaintext)
}

func TripleDES128_Encrypt_CFB8(plaintext []byte, key [16]byte, iv [8]byte) (ciphertext []byte, err error) {
	plaintext = addPadding(plaintext)
	ciphertext = make([]byte, len(plaintext))
	ciphertext[0] = TripleDES128_Encrypt64(iv, key)[0] ^ plaintext[0]
	for i := 1; i < len(plaintext); i++ {
		for j := 0; j < 7; j++ {
			iv[j] = iv[j+1]
		}
		iv[7] = ciphertext[i-1]
		ciphertext[i] = TripleDES128_Encrypt64(iv, key)[0] ^ plaintext[i]
	}
	return ciphertext, nil
}

func TripleDES128_Decrypt_CFB8(ciphertext []byte, key [16]byte, iv [8]byte) (plaintext []byte, err error) {
	if len(ciphertext)%8 != 0 {
		return plaintext, errors.New("invalid input ciphertext")
	}
	plaintext = make([]byte, len(ciphertext))
	plaintext[0] = TripleDES128_Encrypt64(iv, key)[0] ^ ciphertext[0]
	for i := 1; i < len(ciphertext); i++ {
		for j := 0; j < 7; j++ {
			iv[j] = iv[j+1]
		}
		iv[7] = ciphertext[i-1]
		plaintext[i] = TripleDES128_Encrypt64(iv, key)[0] ^ ciphertext[i]
	}
	return removePadding(plaintext)
}

func TripleDES128_Encrypt_OFB(plaintext []byte, key [16]byte, nonce [8]byte) (ciphertext []byte, err error) {
	plaintext = addPadding(plaintext)
	ciphertext = make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += 8 {
		nonce = TripleDES128_Encrypt64(nonce, key)
		for j := 0; j < 8; j++ {
			ciphertext[i+j] = nonce[j] ^ plaintext[i+j]
		}
	}
	return ciphertext, nil
}

func TripleDES128_Decrypt_OFB(ciphertext []byte, key [16]byte, nonce [8]byte) (plaintext []byte, err error) {
	plaintext = make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += 8 {
		nonce = TripleDES128_Encrypt64(nonce, key)
		for j := 0; j < 8; j++ {
			plaintext[i+j] = nonce[j] ^ ciphertext[i+j]
		}
	}
	return removePadding(plaintext)
}

func TripleDES192_Encrypt_CBC(plaintext []byte, key [24]byte, iv [8]byte) (ciphertext []byte) {
	plaintext = addPadding(plaintext)
	ciphertext = make([]byte, 0, len(plaintext))
	var block [8]byte
	for i := 0; i < len(plaintext); i += 8 {
		block = xor64bit(([8]byte)(plaintext[i:i+8]), iv)
		block = TripleDES192_Encrypt64(block, key)
		iv = block
		ciphertext = append(ciphertext, block[:]...)
	}
	return ciphertext
}

func TripleDES192_Decrypt_CBC(ciphertext []byte, key [24]byte, iv [8]byte) (plaintext []byte, err error) {
	plaintext = make([]byte, 0, len(ciphertext))
	var block [8]byte
	var ciphertextblock [8]byte
	for i := 0; i < len(ciphertext); i += 8 {
		ciphertextblock = ([8]byte)(ciphertext[i : i+8])
		block = xor64bit(iv, TripleDES192_Decrypt64(ciphertextblock, key))
		iv = ciphertextblock
		plaintext = append(plaintext, block[:]...)
	}
	return removePadding(plaintext)
}

func TripleDES192_Encrypt_CFB8(plaintext []byte, key [24]byte, iv [8]byte) (ciphertext []byte, err error) {
	plaintext = addPadding(plaintext)
	ciphertext = make([]byte, len(plaintext))
	ciphertext[0] = TripleDES192_Encrypt64(iv, key)[0] ^ plaintext[0]
	for i := 1; i < len(plaintext); i++ {
		for j := 0; j < 7; j++ {
			iv[j] = iv[j+1]
		}
		iv[7] = ciphertext[i-1]
		ciphertext[i] = TripleDES192_Encrypt64(iv, key)[0] ^ plaintext[i]
	}
	return ciphertext, nil
}

func TripleDES192_Decrypt_CFB8(ciphertext []byte, key [24]byte, iv [8]byte) (plaintext []byte, err error) {
	if len(ciphertext)%8 != 0 {
		return plaintext, errors.New("invalid input ciphertext")
	}
	plaintext = make([]byte, len(ciphertext))
	plaintext[0] = TripleDES192_Encrypt64(iv, key)[0] ^ ciphertext[0]
	for i := 1; i < len(ciphertext); i++ {
		for j := 0; j < 7; j++ {
			iv[j] = iv[j+1]
		}
		iv[7] = ciphertext[i-1]
		plaintext[i] = TripleDES192_Encrypt64(iv, key)[0] ^ ciphertext[i]
	}
	return removePadding(plaintext)
}

func TripleDES192_Encrypt_OFB(plaintext []byte, key [24]byte, nonce [8]byte) (ciphertext []byte, err error) {
	plaintext = addPadding(plaintext)
	ciphertext = make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += 8 {
		nonce = TripleDES192_Encrypt64(nonce, key)
		for j := 0; j < 8; j++ {
			ciphertext[i+j] = nonce[j] ^ plaintext[i+j]
		}
	}
	return ciphertext, nil
}

func TripleDES192_Decrypt_OFB(ciphertext []byte, key [24]byte, nonce [8]byte) (plaintext []byte, err error) {
	plaintext = make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += 8 {
		nonce = TripleDES192_Encrypt64(nonce, key)
		for j := 0; j < 8; j++ {
			plaintext[i+j] = nonce[j] ^ ciphertext[i+j]
		}
	}
	return removePadding(plaintext)
}
