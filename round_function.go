package godes

// This file defines the round function f and its components

// expand function takes 32-bit input and expands it to 48-bits
func expand(input [4]byte) (output [6]byte) {
	var temp byte = 0
	for ind, val := range EXPANSION {
		if val%8 != 0 {
			temp = (temp << 1) | ((input[(val-1)/8] >> (8 - val%8)) & 1)
		} else {
			temp = (temp << 1) | (input[(val-1)/8] & 1)
		}
		output[ind/8] = temp
	}
	return output
}
