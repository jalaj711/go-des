# GO-DES

A Golang implementation of the Data Encryption Standard (DES). This module provides mainly two functions: `Encrypt64` and `Decrypt(64)`.

## Encrypt64

The Encrypt64 function signature is:

```golang
func Encrypt64(data [8]byte, key [8]byte) (encrypted [8]byte)
```

It takes a 64-bit data block and a 64-bit key and returns an encrypted version of the plaintext which is also 64-bits. All the inputs and outputs are represented as byte arrays of fixed size 8.

## Decrypt64

The Decrypt64 function signature is:

```golang
func Decrypt64(data [8]byte, key [8]byte) (decrypted [8]byte)
```

This is the inverse of the `Encrypt64` function. It takes in a 64-bit ciphertext and a 64-bit key to generate a 64 bit plaintext version of it.
