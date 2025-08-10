package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zrygan/symbolang/symerr"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		fname := args[1]

		if !strings.HasSuffix(fname, ".🤓") {
			fmt.Println("Symbolab Error: ")
			os.Exit(1)
			symerr.ErrorMessage(
				"File provided is not a symbolang source file.",
				"",
				&symerr.ErrorType{ReadError: true},
			)
		}

	} else if len(args) > 2 {
		symerr.ErrorMessage(
			"Only provide one source file as input.",
			"Use `go run main.go <fileName>`.",
			&symerr.ErrorType{ReadError: true},
		)
	} else {
		symerr.ErrorMessage(
			"No source file provided.",
			"Use `go run main.go <fileName>`.",
			&symerr.ErrorType{ReadError: true},
		)
	}

	// If everything is fine, you can continue normally or exit with 0
	os.Exit(0)
}
