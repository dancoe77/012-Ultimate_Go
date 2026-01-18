package nlp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	require.New(t)
	// setup: call a function
	// teardown: defer/t.Cleanup
	text := "Who's on first?"
	tokens := Tokenize(text)
	expected := []string{"who", "s", "on", "first"}
	require.Equal(t, expected, tokens)
	/* Before testify
	if !slices.Equal(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
	*/
}

func TestTokenizeTable(t *testing.T) {
	var cases = []struct {
		text   string
		tokens []string
	}{
		{"Who's on first?", []string{"who", "s", "on", "first"}},
		{"What's on second?", []string{"what", "s", "on", "second"}},
		{"", nil},
	}
	for _, tc := range cases {
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			require.Equal(t, tc.tokens, tokens)
			/* Before testify
			if !slices.Equal(tc.tokens, tokens) {
				t.Fatalf("expected %#v, got %#v", tc.tokens, tokens)
			}

			*/
		})
	}
}

/*
Selecting tests
- "-run" flag: regexp
- build tags (//go:build comment)
- environment variables
*/

// In Jenkins use BUILD_NUMBER
var inCI = os.Getenv("CI") != ""

func TestInCI(t *testing.T) {
	if !inCI {
		t.Skip("not in CI")
	}
}

// CI=yes go test -v otherwise TestInCI will skip
