package strutil

import (
	"testing"
)

// TestSubstring tests Substring for returing correct values on various input.
func TestSubstring(t *testing.T) {
	t.Parallel()

	type Case struct {
		Value, Result       string
		Start, End          int
		TrimSpace, Ellipsis bool
	}

	cases := []Case{
		{
			Value:     "abcdef",
			Start:     0,
			End:       2,
			TrimSpace: false,
			Ellipsis:  false,
			Result:    "ab",
		},
		{
			Value:     "abcdef",
			Start:     4,
			End:       9,
			TrimSpace: false,
			Ellipsis:  false,
			Result:    "ef",
		},
		{
			Value:     "abcdef",
			Start:     1,
			End:       3,
			TrimSpace: false,
			Ellipsis:  false,
			Result:    "bc",
		},
		// ellipsis: true
		{
			Value:     "abcdef",
			Start:     0,
			End:       2,
			TrimSpace: false,
			Ellipsis:  true,
			Result:    "ab...",
		},
		{
			Value:     "abcdef",
			Start:     4,
			End:       9,
			TrimSpace: false,
			Ellipsis:  true,
			Result:    "ef",
		},
		{
			Value:     "abcdef",
			Start:     1,
			End:       3,
			TrimSpace: false,
			Ellipsis:  true,
			Result:    "bc...",
		},
		// trimSpace: true
		{
