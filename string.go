package strutil

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Substring cuts a string from start to end optionally add ellipsis if
// there is still more characters remaining in the string.
func Substring(s string, start int, end int, trimSpace, ellipsis bool) string {
	if end < 0 {
		return s
	}

	if trimSpace {
		s = strings.TrimSpace(s)
	}

	var (
		index = 0
		i     = 0
	)

	const tail = "..."

	for j := range s {
		if i == start {
			index = j
		}

		if i == end {
			if len(s) > j && ellipsis {
				return s[index:j] + tail
			}

			return s[index:j]
		}
		i++
	}

	return s[index:]
}

// TrimPhrases trims phrases prefix/suffix from s.
func TrimPhrases(s string, trimSpace bool, phrases ...string) string {
	s = TrimPrefixPhrases(s, trimSpace, phrases...)
	s = TrimSuffixPhrases(s, trimSpace, phrases...)

	return s
}

// TrimSuffixPhrases trims phrases suffix from s.
func TrimSuffixPhrases(s string, trimSpace bool, phrases ...string) string {
	for _, v := range phrases {
		if v != "" {
			if trimSpace {
				s = strings.TrimSpace(s)
			}

			for strings.HasSuffix(s, v) {
				s = strings.TrimSuffix(s, v)
				if trimSpace {
					s = strings.TrimSpace(s)
				}
			}
		}
	}

	return s
}

// TrimPrefixPhrases trims phrases prefix from s.
func TrimPrefixPhrases(s string, trimSpace bool, phrases ...string) string {
	for _, str := range phrases {
		if str != "" {
			if trimSpace {
				s = strings.TrimSpace(s)
			}

			for strings.HasPrefix(s, str) {
				s = strings.TrimPrefix(s, str)
				if trimSpace {
					s = strings.TrimSpace(s)
				}
			}
		}
	}

	return s
}

// IsNumeric returns true if s contains only utf8 numbers.
func IsNumeric(s string, trimSpace bool) bool {
	if trimSpace {
		s = strings.TrimSpace(s)
	}

	for _, char := range strings.Split(s, "") {
		r, _ := utf8.DecodeRuneInString(char)
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

// IsMultiByte checks if s contains a multibyte characters.
func IsMultiByte(s string, trimSpace bool) bool {
	if trimSpace {
		s = strings.TrimSpace(s)
	}

	for _, char := range strings.Split(s, "") {
		_, v := utf8.DecodeRuneInString(char)
		if v > 1 {
			return true
		}
	}

	return false
}

// WrappedASCII wraps s of ASCII long lines to max per line.
func WrappedASCII(s string, max int) (lines []string) {
	if max == 0 {
		return
	}

	l := len(s)

	end := l - 1
	if max >= end {
		lines = append(lines, s)

		return
	}

	var i, index int

	for j := range s {
		if j > 0 && i%max == 0 {
			lines = append(lines, s[index:j])
			index = j
		}

		if j == end {
			lines = append(lines, s[index:])
		}

		i++
	}

	return
}

// AsString converts any value(interface) to string.
func AsString(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

// ToCSV converts []interface{} to CSV value.
func ToCSV(v []interface{}, glue string) string {
	const sep = ","

	elems := []string{}
	for _, i := range v {
		elems = append(elems, AsString(i))
	}

	if glue != "" {
		return strings.Join(elems, glue)
	}

	return strings.Join(elems, sep)
}

// ToSlice converts interface to []string.
func ToSlice(v interface{}) []string {
	if s, ok := v.([]string); ok {
		return s
	}

	return []string{
		fmt.Sprintf("%v", v),
	}
}

// SplitTrim splits v to []string  using by.
func SplitTrim(v, by string) []string {
	s := strings.Split(v, by)
