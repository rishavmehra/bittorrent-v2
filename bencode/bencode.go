package bencode

import (
	"fmt"
	"strconv"
	"unicode"
)

// 5:hello => hello
// 6:rishav => rishav
func DecodeBencodeString(bencodedString string) (string, int, error) {
	index := 0

	switch {
	case unicode.IsDigit(rune(bencodedString[0])):
		var colonIndex int

		// this give us colon index
		for i := 0; i < len(bencodedString); i++ {
			if bencodedString[i] == ':' {
				colonIndex = i
				break
			}
		}

		// string(length number of the string) before the colon
		lengthStr := bencodedString[:colonIndex]

		// convert the length number of string ASCII to integer
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return "", index, err
		}

		index = colonIndex + 1 + length
		decodedString := bencodedString[colonIndex+1 : index]
		return decodedString, index, nil
	default:
		return "", 0, fmt.Errorf("invalid BecodeString %q", bencodedString)
	}
}

func decodeBencodeInteger(bencodedString string) (int, int, error) {
	index := 0

}
