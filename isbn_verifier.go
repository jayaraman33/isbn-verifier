package isbn

import (
	"unicode"
)

//IsValidISBN is the verification function
func IsValidISBN(s string) bool {
	n := 10
	var sum int
	for _, v := range s {
		if unicode.IsNumber(v) {
			sum += n * int(v-'0')
			n--
		} else if n == 1 && v == 'X' {
			sum += n * 10
			n--
		}
	}
	return n == 0 && sum%11 == 0
}
