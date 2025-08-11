package parser

import (
	"fmt"

	"github.com/zrygan/symbolang/tokens"
)

type Statement interface {
	Execute()
}

type PrintStatement struct {
	PrintType tokens.TokenType
	Value     string
}

func (ps *PrintStatement) Execute() {
	switch ps.PrintType {
	case tokens.Pencil:
		fmt.Print(ps.Value)
	case tokens.Pen:
		fmt.Println(ps.Value)
	}
}
