package godes

func getReducedKey(key [8]byte) [7]byte {
	reduced := [7]byte{}
	var temp byte = 0
	for ind, val := range KEY_PERMUTATION_CHOICE_1 {
		temp = (temp << 1) | ((key[val/8] >> (8 - val%8)) & 1)
		reduced[ind/8] = temp
	}
	return reduced
}

// leftRotateKey takes a 56-bit key in the format of a byte array by rotationAmount
// It does so in 2 groups, ie the key is divided into two groups of 28-bits each
// and both groups are rotated seperately
func leftRotateKey(key [7]byte, rotationAmount uint8) (rotatedKey [7]byte) {

	// stores the first bit(s) of input to be concatenated to the last of shifted bits
	lastOutputBits := (key[0] >> (8 - rotationAmount))

	var temp byte
	var i byte

	// stores as many ones as the rotationAmount
	// used to extract bits from a byte
	var allOnes byte = 0
	for i = rotationAmount; i > 0; i-- {
		allOnes = (allOnes << 1) | 1
	}

	// left shifts bit 1 to bit (28-rotationAmount)
	for i = 0; i < 28-rotationAmount; i++ {
		temp = (temp << 1) | ((key[(i+rotationAmount)/8] >> (7 - (i+rotationAmount)%8)) & 1)
		rotatedKey[i/8] = temp
	}
	// sets last bits of this rotation
	temp = (temp << rotationAmount) | lastOutputBits
	rotatedKey[3] = temp

	// update output bits for second half
	lastOutputBits = (key[3] >> (4 - rotationAmount)) & allOnes
	for i = 28; i < 56-rotationAmount; i++ {
		temp = (temp << 1) | ((key[(i+rotationAmount)/8] >> (7 - (i+rotationAmount)%8)) & 1)
		rotatedKey[i/8] = temp
	}
	temp = (temp << rotationAmount) | lastOutputBits
	rotatedKey[6] = temp

	return rotatedKey
}

// getRotatedRoundKeys takes the initial 56-bit key and generates 16 56-bit left-rotated keys for each round
// the keys are shifted according to the key schedule defined in tables.go
func getRotatedRoundKeys(key [7]byte) (rotatedRoundKeys [16][7]byte) {
	rotatedRoundKeys[0] = leftRotateKey(key, KEY_ROTATION[0])
	for i := 1; i < 16; i++ {
		rotatedRoundKeys[i] = leftRotateKey(rotatedRoundKeys[i-1], KEY_ROTATION[i])
	}
	return rotatedRoundKeys
}

func getPermutedRoundKeys(rotatedRoundKeys [16][7]byte) (permutedRoundKeys [16][6]byte) {
	reduced := [6]byte{}
	var temp byte = 0
	for key_ind, key := range rotatedRoundKeys {
		for ind, val := range KEY_PERMUTATION_CHOICE_2 {
			if val%8 != 0 {
				temp = (temp << 1) | ((key[(val-1)/8] >> (8 - val%8)) & 1)
			} else {
				temp = (temp << 1) | (key[(val-1)/8] & 1)
			}
			reduced[ind/8] = temp
		}
		permutedRoundKeys[key_ind] = reduced
	}
	return permutedRoundKeys
}
