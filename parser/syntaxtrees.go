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
	output, exists := globalEnv[ps.Value]
	if !exists {
		// treat as literal
		output = ps.Value
	}

	switch ps.PrintType {
	case tokens.Pencil:
		fmt.Print(output)
	case tokens.Pen:
		fmt.Println(output)
	}
}

type VariableStatement struct {
	Name  string
	Value interface{}
}

func (vs *VariableStatement) Execute() {
	globalEnv[vs.Name] = vs.Value
}
