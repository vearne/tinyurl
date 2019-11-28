package libs

import "errors"

const ALL_CHAR = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func UintToBase62(value uint64) string {
	var temp []byte = make([]byte, 0, 10)
	for value > 0 {
		temp = append(temp, ALL_CHAR[value%62])
		value = value / 62
	}
	reverse(temp)
	return string(temp)
}

func Base62ToUint(str string) (uint64, error) {
	var value uint64 = 0
	var ch byte
	var idx uint64

	for i := 0; i < len(str); i++ {
		ch = str[i]
		if ch >= '0' && ch <= '9' {
			idx = uint64(ch - '0')
		} else if ch >= 'A' && ch <= 'Z' {
			idx = uint64(ch - 'A' + 10)
		} else if ch >= 'a' && ch <= 'z' {
			idx = uint64(ch - 'a' + 36)
		} else {
			return 0, errors.New("invalid byte")
		}
		value = value*62 + idx
	}
	return value, nil
}

func reverse(s []byte){
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}
