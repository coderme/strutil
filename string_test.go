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
			Value:     " abcdef ",
			Start:     0,
			End:       2,
			TrimSpace: true,
			Ellipsis:  false,
			Result:    "ab",
		},
		{
			Value:     " abcdef ",
			Start:     4,
			End:       9,
			TrimSpace: true,
			Ellipsis:  false,
			Result:    "ef",
		},
		{
			Value:     " abcdef ",
			Start:     1,
			End:       3,
			TrimSpace: true,
			Ellipsis:  false,
			Result:    "bc",
		},

		// trimSpace: true
		// ellipsis: true
		{
			Value:     " abcdef ",
			Start:     0,
			End:       2,
			TrimSpace: true,
			Ellipsis:  true,
			Result:    "ab...",
		},
		{
			Value:     " abcdef ",
			Start:     4,
			End:       9,
			TrimSpace: true,
			Ellipsis:  true,
			Result:    "ef",
		},
		{
			Value:     " abcdef ",
			Start:     1,
			End:       3,
			TrimSpace: true,
			Ellipsis:  true,
			Result:    "bc...",
		},
	}

	for _, v := range cases {
		if r := Substring(v.Value, v.Start, v.End, v.TrimSpace, v.Ellipsis); r != v.Result {
			t.Error("expected:", v.Result,
				", but found:", r,
				"for:", v.Value,
				"Start:", v.Start,
				"end:", v.End,
				"trimSpace:", v.TrimSpace,
				"ellipsis:", v.Ellipsis,
			)
		}
	}
}

// TestTrimPhrases tests TrimPhrases for trimming phrases from the both start and end of a string.
func TestTrimPhrases(t *testing.T) {
	t.Parallel()

	type Case struct {
		Value, Result string
		Phrases       []string
		TrimSpace     bool
	}

	cases := []Case{
		{
			Value: "I bought a new car which is really expensive",

			Phrases: []string{
				"I bought a",
				"really expensive",
			},

			Result:    "new car which is",
			TrimSpace: true,
		},
		{
			Value: "I bought a new car which is really expensive",
