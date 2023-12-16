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
