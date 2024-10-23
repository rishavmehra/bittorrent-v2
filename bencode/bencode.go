package bencode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// 5:hello => hello
// 6:rishav => rishav
// <string length encoded in base ten ASCII>:<string data>
func DecodeBencodeString(bencodedString string) (string, int, error) {
	index := 0

	switch {
	case unicode.IsDigit(rune(bencodedString[index])):
		var firstColonIndex int

		for i := 0; i < len(bencodedString); i++ {
			if bencodedString[i] == ':' {
				firstColonIndex = i
				break
			}
		}

		lengthStr := bencodedString[:firstColonIndex]

		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return "", index, err
		}

		index = firstColonIndex + 1 + length
		decodedString := bencodedString[firstColonIndex+1 : index]

		return decodedString, index, nil
	default:
		return "", 0, fmt.Errorf("invalid BencodeString %q", bencodedString)
	}
}

// i3e => 3
// i-3e => -3
// All encodings with a leading zero, such as i03e, are invalid
// i<integer encoded in base ten ASCII>e
func DecodeBencodeInteger(bencodedString string) (int, int, error) {
	index := 0

	switch c := rune(bencodedString[index]); c {
	case 'i':
		// strings.Index("Rishav", "hav") => 3
		indexEnd := strings.Index(bencodedString, "e")
		integer, err := strconv.Atoi(bencodedString[1:indexEnd])
		if err != nil {
			return 0, 0, err
		}

		index = indexEnd + 1
		return integer, index, nil
	default:
		return 0, 0, fmt.Errorf("invalid BecodeInteger %q", bencodedString)
	}
}

/*
l
 4:spam
 4:eggs
e
represents the list of two strings: [ "spam", "eggs" ]
*/
// l<bencoded values>e
func DecodeBencodeList(bencodedString string) ([]interface{}, int, error) {
	index := 0
	depth := 0
	l := make([]interface{}, 0)

	for {
		switch c := rune(bencodedString[index]); {
		case c == 'e':
			return l, index + 1, nil
		case c == 'l':
			if depth == 0 {
				index += 1
				depth += 1
			} else {
				nl, relIndex, err := DecodeBencodeList(bencodedString[index:])
				if err != nil {
					return nil, index, err
				}
				l = append(l, nl...)
				index += relIndex
			}
		case c == 'i':
			i, relIndex, err := DecodeBencodeInteger(bencodedString[index:])
			if err != nil {
				return nil, index, err
			}
			l = append(l, i)
			index += relIndex
		case unicode.IsDigit(c):
			s, relIndex, err := DecodeBencodeString(bencodedString[index:])
			if err != nil {
				return nil, index, err
			}
			l = append(l, s)
			index += relIndex
		default:
			return nil, index, fmt.Errorf("unknow:\n%s\n", bencodedString[index:])
		}
	}
}

/*
d
 4:spam
  l
   1:a
   1:b
  e
e represents the dictionary { "spam" => [ "a", "b" ] }
*/
// d<bencoded string><bencoded element>e
func DecodeBencodeDic(bencodedString string) (map[string]interface{}, int, error) {
	index := 0
	depth := 0
	isValue := false
	key := ""
	d := make(map[string]interface{})

	for {
		switch c := rune(bencodedString[index]); {
		case c == 'e':
			return d, index + 1, nil
		case c == 'd':
			switch {
			case depth == 0:
				index += 1
				depth += 1
			case isValue:
				// TODO: refactor recursion out
				nd, relIndex, err := DecodeBencodeDic(bencodedString[index:])
				if err != nil {
					return nil, index, err
				}
				d[key] = nd
				index += relIndex
				isValue = false
			default:
				return nil, index, fmt.Errorf("invalid BencodeDict %q", bencodedString[index:])
			}
		case unicode.IsDigit(c):
			switch {
			case depth == 0:
				return nil, index, fmt.Errorf("invalid BencodeDict %q", bencodedString[index:])
			case !isValue:
				s, relIndex, err := DecodeBencodeString(bencodedString[index:])
				if err != nil {
					return nil, index, err
				}
				key = s
				index += relIndex
				isValue = true
			case isValue:
				s, relIndex, err := DecodeBencodeString(bencodedString[index:])
				if err != nil {
					return nil, index, err
				}
				d[key] = s
				index += relIndex
				isValue = false
			default:
				return nil, index, fmt.Errorf("invalid BencodeDict, unexpected string %q", bencodedString[index:])
			}
		case c == 'i':
			if isValue {
				i, relIndex, err := DecodeBencodeInteger(bencodedString[index:])
				if err != nil {
					return nil, index, err
				}
				d[key] = i
				index += relIndex
				isValue = false
			} else {
				return nil, index, fmt.Errorf("invalid BencodeDict, unexpected integer %q", bencodedString[index:])
			}
		case c == 'l':
			if isValue {
				l, relIndex, err := DecodeBencodeList(bencodedString[index:])
				if err != nil {
					return nil, index, err
				}
				d[key] = l
				index += relIndex
				isValue = false
			} else {
				return nil, index, fmt.Errorf("invalid BencodeDict, unexpected list %q", bencodedString[index:])
			}
		default:
			return nil, index, fmt.Errorf("invalid BencodeDict %q", bencodedString[index:])
		}
	}
}

func DecodeBencode(bencodedString string) (interface{}, error) {
	c := rune(bencodedString[0])
	switch {
	case unicode.IsDigit(c):
		result, _, err := DecodeBencodeString(bencodedString)
		if err != nil {
			return "", err
		}
		return result, nil
	case c == 'i':
		result, _, err := DecodeBencodeInteger(bencodedString)
		if err != nil {
			return "", err
		}
		return result, nil
	case c == 'l':
		result, _, err := DecodeBencodeList(bencodedString)
		if err != nil {
			return "", err
		}
		return result, nil
	case c == 'd':
		result, _, err := DecodeBencodeDic(bencodedString)
		if err != nil {
			return "", err
		}
		return result, nil
	default:
		return "", fmt.Errorf("unSupported:\n%s\n", bencodedString)
	}
}
