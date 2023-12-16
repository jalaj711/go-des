package godes

// This file includes functions for initial permutation and inverse initial permutation

// initial_permutation takes a 64-bit plaintext and produces a 64-bit permuted version of it
// as defined by the INITIAL_PERMUTATION table in tables.go
func initial_permutation(plaintext [8]byte) (permuted [8]byte) {
	var temp byte = 0
	for ind, val := range INITIAL_PERMUTATION {
		if val%8 != 0 {
			temp = (temp << 1) | ((plaintext[(val-1)/8] >> (8 - val%8)) & 1)
		} else {
			temp = (temp << 1) | (plaintext[(val-1)/8] & 1)
		}
		permuted[ind/8] = temp
	}
	return permuted
}

// inverse_initial_permutation takes a 64-bit permuted value and produces a 64-bit plaintext version of it
// as defined by the INVERSE_INITIAL_PERMUTATION table in tables.go
func inverse_initial_permutation(permuted [8]byte) (plaintext [8]byte) {
	var temp byte = 0
	for ind, val := range INVERSE_INITIAL_PERMUTATION {
		if val%8 != 0 {
			temp = (temp << 1) | ((permuted[(val-1)/8] >> (8 - val%8)) & 1)
		} else {
			temp = (temp << 1) | (permuted[(val-1)/8] & 1)
		}
		plaintext[ind/8] = temp
	}
	return plaintext
}
