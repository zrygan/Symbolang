package main

import (
	"os"
	"strings"

	"github.com/zrygan/symbolang/lexer"
	"github.com/zrygan/symbolang/options"
	"github.com/zrygan/symbolang/parser"
	"github.com/zrygan/symbolang/symerr"
)

func main() {
	args := os.Args

	if len(args) >= 2 {
		fname := args[1]
		opts := args[2:]

		if !strings.HasSuffix(fname, ".🤓") {
			symerr.ErrorMessage(
				"File provided is not a symbolang source file.",
				"",
				&symerr.ErrorType{FatalErr: true},
			)
		}

		options.Options(opts)

		s := lexer.OpenFile(fname)
		l := lexer.NewLex(s.SourceFile)
		p := parser.NewParser(l)

		program := p.ParseProgram()

		for _, stmt := range program {
			stmt.Execute()
		}

	} else {
		symerr.ErrorMessage(
			"No source file provided.",
			"Use `go run main.go <fileName>`.",
			&symerr.ErrorType{FatalErr: true},
		)
	}

	os.Exit(0)
}
