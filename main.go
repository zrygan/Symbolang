package main

import (
	"os"
	"strings"

	"github.com/zrygan/symbolang/interpreter"
	"github.com/zrygan/symbolang/symerr"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		fname := args[1]

		if !strings.HasSuffix(fname, ".🤓") {
			symerr.ErrorMessage(
				"File provided is not a symbolang source file.",
				"",
				&symerr.ErrorType{FatalErr: true},
			)
		}

		source := interpreter.OpenFile(fname)
		source.DebugSource()
	} else if len(args) > 2 {
		symerr.ErrorMessage(
			"Only provide one source file as input.",
			"Use `go run main.go <fileName>`.",
			&symerr.ErrorType{FatalErr: true},
		)
	} else {
		symerr.ErrorMessage(
			"No source file provided.",
			"Use `go run main.go <fileName>`.",
			&symerr.ErrorType{FatalErr: true},
		)
	}

	// If everything is fine, you can continue normally or exit with 0
	os.Exit(0)
}
