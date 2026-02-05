package nlp

import (
	"os"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	require.New(t)
	// setup: call a function
	// teardown: defer/t.Cleanup
	text := "Who's on first?"
	tokens := Tokenize(text)
	expected := []string{"who", "on", "first"}
	require.Equal(t, expected, tokens)
	/* Before testify
	if !slices.Equal(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
	*/
}

type tokCase struct {
	Text   string
	Tokens []string
	Name   string
}

func loadTokenizeCases(t *testing.T) []tokCase {
	// t.Helper()
	file, err := os.Open("testdata/tokenize_cases.toml")
	require.NoError(t, err)
	defer file.Close() // t.Cleanup(file.Close)

	var data struct {
		Cases []tokCase `toml:"case"`
	}
	dec := toml.NewDecoder(file)
	_, err = dec.Decode(&data)
	require.NoError(t, err)
	return data.Cases
}

// Exercise: Read test cases from tokenize_cases.toml
// instead of in-memory slice
func TestTokenizeTable(t *testing.T) {
	/*
		var cases = []struct {
			text   string
			tokens []string
		}{
			{"Who's on first?", []string{"who", "s", "on", "first"}},
			{"What's on second?", []string{"what", "s", "on", "second"}},
			{"", nil},
		}
	*/
	cases := loadTokenizeCases(t)

	for _, tc := range cases {
		name := tc.Name
		if name == "" {
			name = tc.Text
		}
		t.Run(name, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			// TOML does not have nil
			if tokens == nil {
				tokens = []string{}
			}
			require.Equal(t, tc.Tokens, tokens)
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

func FuzzTokenize(f *testing.F) {
	f.Add("")
	fn := func(t *testing.T, text string) {
		tokens := Tokenize(text)
		lText := strings.ToLower(text)
		for _, tok := range tokens {
			require.Contains(t, lText, tok) // + "XXX" = failed test
		}
	}
	f.Fuzz(fn)
}
