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
