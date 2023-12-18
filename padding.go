package godes

// addPadding adds padding to the input to make it a multiple of 64-bits
// it works according to the padding defined in PKCS#7 (RFC5652)
func addPadding(input []byte) []byte {
	size := len(input)
	paddingBytes := 8 - size%8
	padding := make([]byte, paddingBytes)
	for i := range padding {
		padding[i] = byte(paddingBytes)
	}
	return append(input, padding...)
}

// removePadding removes padding from the input to tryutn the original text
// it works according to the padding defined in PKCS#7 (RFC5652)
func removePadding(input []byte) []byte {
	return input[:len(input)-int(input[len(input)-1])]
}
