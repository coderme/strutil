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
