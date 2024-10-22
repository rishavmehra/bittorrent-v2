package bencode

import (
	"fmt"
	"strconv"
	"unicode"
)

func DecodeBencode(bencodedString string) (interface{}, error) {
	if unicode.IsDigit(rune(bencodedString[0])) {

		var colonIndex int

		for i := 0; i < len(bencodedString); i++ {
			if bencodedString[i] == ':' {
				colonIndex = i
				break
			}
		}

		lengthStr := bencodedString[:colonIndex]
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return "", err
		}

		return bencodedString[colonIndex+1 : colonIndex+1+length], nil
	} else if rune(bencodedString[0]) == 'i' {
		return strconv.Atoi(bencodedString[1 : len(bencodedString)-1])
	} else {
		return "", fmt.Errorf("only string are supported at the moment")
	}
}
