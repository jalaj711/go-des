# GO-DES

A Golang implementation of the Data Encryption Standard (DES). This module provides the following functions:

## Block level

### Encrypt64

The Encrypt64 function signature is:

```golang
func Encrypt64(data [8]byte, key [8]byte) (encrypted [8]byte)
```

It takes a 64-bit data block and a 64-bit key and returns an encrypted version of the plaintext which is also 64-bits. All the inputs and outputs are represented as byte arrays of fixed size 8.

### Decrypt64

The Decrypt64 function signature is:

```golang
func Decrypt64(data [8]byte, key [8]byte) (decrypted [8]byte)
```

This is the inverse of the `Encrypt64` function. It takes in a 64-bit ciphertext and a 64-bit key to generate a 64 bit plaintext version of it.

## Primary encryption and decryption

### Encrypt

Encrypts using DES in ECB mode.

The Encrypt function signature is:

```golang
func Encrypt(data []byte, key [8]byte) (encrypted []byte)
```

It takes a byte array and a 64-bit key and returns an encrypted version of the plaintext. All the inputs and outputs are represented as byte arrays. The input plaintext is first padded to 64-bits using PKCS7 padding.

### Decrypt

Decrypts using DES in ECB mode.

The Decrypt function signature is:

```golang
func Decrypt(data []byte, key [8]byte) (decrypted []byte, err error)
```

It takes a byte array and a 64-bit key and returns an decrypted version of the ciphertext. All the inputs and outputs are represented as byte arrays. Padding is removed from the decrypted plaintext using PKCS7 padding scheme. Error is raised if input is not a multiple of 64 bits or if padding is invalid.

## Modes of operation

### Encrypt_CBC

Encrypts using DES in CBC mode.

Function signature is

```golang
func Encrypt_CBC(plaintext []byte, key [8]byte, iv [8]byte) (ciphertext []byte)
```

### Decrypt_CBC

Decrypts using DES in CBC mode.

Function signature is

```golang
func Decrypt_CBC(ciphertext []byte, key [8]byte, iv [8]byte) (plaintext []byte, err error)
```

### Encrypt_CFB8

Encrypts using DES in CFB mode with 8-bit transmission.

Function signature is

```golang
func Encrypt_CFB8(plaintext []byte, key [8]byte, iv [8]byte) (ciphertext []byte, err error)
```

### Decrypt_CFB8

Decrypts using DES in CFB mode with 8-bit transmission.

Function signature is

```golang
func Decrypt_CFB8(ciphertext []byte, key [8]byte, iv [8]byte) (plaintext []byte, err error)
```

### Encrypt_OFB

Encrypts using DES in OFB mode.

Function signature is

```golang
func Encrypt_OFB(plaintext []byte, key [8]byte, nonce [8]byte) (ciphertext []byte, err error)
```

### Decrypt_OFB

Decrypts using DES in CFB mode.

Function signature is

```golang
func Decrypt_OFB(ciphertext []byte, key [8]byte, nonce [8]byte) (plaintext []byte, err error)
```

## Triple DES

All the methods defined above are also available in their 3DES versions. Prefixing any of the methods above with `TripleDES128_` will give you it's 2 Key 3DES version. All such methods take 128-bit (16-byte) keys. Similiarly, prefixing with `TripleDES192_` will give you it's 3 Key 3DES version. These methods take 192-bit (24-byte) keys. Besides key-length all functions have the same signature as their normal DES counterparts.
