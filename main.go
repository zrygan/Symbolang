package main

import (
	"os"
	"strings"

	"github.com/zrygan/symbolang/lexer"
	"github.com/zrygan/symbolang/parser"
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

		s := lexer.OpenFile(fname)
		l := lexer.NewLex(s.SourceFile)
		p := parser.NewParser(l)

		program := p.ParseProgram()

		for _, stmt := range program {
			stmt.Execute()
		}

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

	os.Exit(0)
}
