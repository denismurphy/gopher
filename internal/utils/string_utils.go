package utils

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

// ReverseString reverses the input string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func InterpolFormat(template string, args ...interface{}) (str string, err error) {
	var buf bytes.Buffer
	buf.Grow(len(template))

	inBrace := false
	numBuffer := make([]byte, 0, 3) // Pre-allocate for small numbers

	for _, char := range template {
		if char == '{' {
			if inBrace {
				return template, fmt.Errorf("invalid format string: nested braces")
			}
			inBrace = true
			numBuffer = numBuffer[:0] // Reset numBuffer
		} else if char == '}' {
			if !inBrace {
				return template, fmt.Errorf("invalid format string: unmatched closing brace")
			}
			if len(numBuffer) == 0 {
				return template, fmt.Errorf("invalid format string: empty braces")
			}

			num, err := strconv.Atoi(string(numBuffer))
			if err != nil {
				return template, fmt.Errorf("invalid format string: non-numeric value in braces")
			}

			if num < 1 || num > len(args) {
				return template, fmt.Errorf("invalid argument index")
			}

			arg := args[num-1]
			buf.WriteString(toString(arg))

			inBrace = false
		} else if inBrace {
			if !unicode.IsDigit(char) {
				return template, fmt.Errorf("invalid format string: non-digit character in braces")
			}
			numBuffer = append(numBuffer, byte(char))
		} else {
			buf.WriteRune(char)
		}
	}

	if inBrace {
		return template, fmt.Errorf("invalid format string: unclosed brace")
	}

	return buf.String(), nil
}

func toString(arg interface{}) string {
	switch v := arg.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprint(v)
	}
}
