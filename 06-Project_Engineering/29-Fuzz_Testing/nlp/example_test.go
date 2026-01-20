package nlp_test

import (
	"fmt"

	"github.com/dancoe77/nlp"
)

func ExampleTokenize() {
	tokens := nlp.Tokenize("Who's on first?")
	fmt.Println(tokens)

	// Output:
	// [who s on first]
}
