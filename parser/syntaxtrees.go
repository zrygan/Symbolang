package parser

import (
	"fmt"

	"github.com/zrygan/symbolang/symerr"
	"github.com/zrygan/symbolang/tokens"
)

type Statement interface {
	Execute()
}

type PrintStatement struct {
	PrintType tokens.TokenType
	Value     string
	ValueType tokens.TokenType
}

func (ps *PrintStatement) Execute() {
	var output interface{} = ps.Value

	// if it's a literal we don't care (no pre-processing)
	if ps.ValueType == tokens.Identifier {
		val, exists := globalEnv[ps.Value]
		if !exists {
			symerr.ErrorMessage(
				"A ✏️ or 🖊️ statement uses a non-existent idenf fier: "+ps.Value+".",
				"",
				&symerr.ErrorType{VarErr: true},
			)
		}
		output = val
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

type ConstStatement struct {
	Name  string
	Value interface{}
}

func (cs *ConstStatement) Execute() {
	_, ok := globalEnv[cs.Name]

	if !ok {
		globalEnv[cs.Name] = cs.Value
	} else {
		symerr.ErrorMessage(
			"The constant variable "+cs.Name+" is redeclared.",
			"You cannot redeclare constants.",
			&symerr.ErrorType{ConstErr: true},
		)
	}
}

type DeleteStatement struct {
	Name string
}

func (ds *DeleteStatement) Execute() {
	_, ok := globalEnv[ds.Name]

	if ok {
		delete(globalEnv, ds.Name)
	} else {
		symerr.ErrorMessage(
			"A non-existent identifier "+ds.Name+" was tried to be deleted.",
			"You cannot delete non-existent identifiers.",
			&symerr.ErrorType{DelErr: true},
		)
	}
}
