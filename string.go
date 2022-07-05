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
	for c := range s {
		s[c] = strings.TrimSpace(s[c])
	}

	return s
}

// IsIEqual checks if all args are the same regardless of the case characters.
func IsIEqual(trim bool, args ...string) bool {
	if n := len(args); n == 0 {
		return false
	} else if n == 1 {
		return true
	}

	for k, arg := range args {
		if k == 0 {
			continue
		}

		pre := strings.ToLower(args[k-1])
		this := strings.ToLower(arg)

		if trim {
			pre = strings.TrimSpace(pre)
			this = strings.TrimSpace(this)
		}

		if pre != this {
			return false
		}
	}

	return true
}

// IContainsAny checks if any args is in the first arg.
func IContainsAny(s string, args ...string) bool {
	s = strings.ToLower(s)
	for _, a := range args {
		if strings.Contains(s, strings.ToLower(a)) {
			return true
		}
	}

	return false
}

// IcontainsAnyPhrase checks if any args is in the first arg, case insensitive.
func IcontainsAnyPhrase(s string, args ...string) bool {
	s = strings.ToLower(s)
	for _, a := range args {
		if strings.Contains(s, strings.ToLower(a)) {
			return true
		}
	}

	return false
}

// ContainsAnyPhrase checks if any args is in the first arg, case sensitive.
func ContainsAnyPhrase(s string, args ...string) bool {
	for _, a := range args {
		if strings.Contains(s, a) {
			return true
		}
	}

	return false
}

// ContainsAll checks if all args are in the first arg, case sensitive.
func ContainsAll(s string, args ...string) bool {
	for _, a := range args {
		if !strings.Contains(s, a) {
			return false
		}
	}

	return true
}

// ContainsIAll checks if all args are in the first arg, case isensitive.
func ContainsIAll(s string, args ...string) bool {
	s = strings.ToUpper(s)
	for _, a := range args {
		if !strings.Contains(s, strings.ToUpper(a)) {
			return false
		}
	}

	return true
}

// ReplaceAny replaces any args, n times with s.
func ReplaceAny(s, r string, n int, args ...string) string {
	for _, a := range args {
		s = strings.Replace(s, a, r, n)
	}

	return s
}

// CaseToWords splits a word into many based on Case so AboutMe -> About Me.
func CaseToWords(s string, trimSpace bool) string {
	var (
		r = []rune{}
		p rune
	)

	const space rune = 32

	for _, c := range s {
		if unicode.IsUpper(c) &&
			p > 0 &&
			p != 32 {
			r = append(r, space, c)

			continue
		}

		r = append(r, c)
		p = c
	}

	v := string(r)
	if trimSpace {
		return strings.TrimSpace(v)
	}
