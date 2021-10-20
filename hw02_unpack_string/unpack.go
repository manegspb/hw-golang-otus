package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var (
		result   strings.Builder
		prevRune *rune
		esc      = false
	)

	for _, r := range input {
		if r == '\\' && !esc {
			esc = true
			continue
		}

		if r == 32 { // "32". This is 'space' in unicode
			return "", ErrInvalidString
		}

		if unicode.IsDigit(r) && !esc {
			if prevRune == nil {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(string(r))
			result.WriteString(strings.Repeat(string(*prevRune), count))
			prevRune = nil
		} else {
			if prevRune != nil {
				result.WriteRune(*prevRune)
			} else {
				prevRune = new(rune)
			}

			*prevRune = r
			esc = false
		}
	}

	if prevRune != nil {
		result.WriteRune(*prevRune)
	}

	return result.String(), nil
}
