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
