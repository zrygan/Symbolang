package parser

import (
	"github.com/zrygan/symbolang/lexer"
	"github.com/zrygan/symbolang/symerr"
	"github.com/zrygan/symbolang/tokens"
)

type Parser struct {
	Lexer   *lexer.Lex
	CurTok  tokens.Token
	peekTok tokens.Token
}

func NewParser(l *lexer.Lex) *Parser {
	p := &Parser{Lexer: l}
	p.peekTok = p.Lexer.ReadNextToken()
	p.nextToken()
	return p
}

func (p *Parser) ParseProgram() []Statement {
	statements := []Statement{}

	for p.CurTok.Type != tokens.EOF {
		// symerr.DebugMessage(
		// 	"ParseProgram",
		// 	p.CurTok.Type,
		// 	p.CurTok.Literal,
		// )

		switch p.CurTok.Type {
		case tokens.Pencil, tokens.Pen:
			stmt := p.ParsePrintStatement()
			statements = append(statements, stmt)
		case tokens.Paper:
			stmt := p.ParseVariable()
			statements = append(statements, stmt)
		case tokens.Rock:
			stmt := p.ParseConst()
			statements = append(statements, stmt)
		case tokens.Scissor:
			stmt := p.ParseDelete()
			statements = append(statements, stmt)
		default:
			symerr.ErrorMessage(
				"Source code includes an unrecognized emoji.",
				"Only include valid emojis.",
				&symerr.ErrorType{FatalErr: true},
			)

			p.nextToken()
		}
	}

	return statements
}

func (p *Parser) ParsePrintStatement() *PrintStatement {
	stmt := &PrintStatement{
		PrintType: p.CurTok.Type,
	}

	// consume print emoji
	p.nextToken()

	if p.CurTok.Type == tokens.Literal || p.CurTok.Type == tokens.Identifier {
		stmt.Value = p.CurTok.Literal
	} else {
		symerr.ErrorMessage(
			"A ✏️ or 🖊️ statement must have a literal or identifier input.",
			"",
			&symerr.ErrorType{FatalErr: true},
		)
	}

	p.checkForStop()

	return stmt
}

func (p *Parser) ParseVariable() *VariableStatement {
	stmt := &VariableStatement{}

	// Move declare emoji
	p.nextToken()

	// Expect a variable name (identifier)
	if p.CurTok.Type != tokens.Identifier {
		symerr.ErrorMessage(
			"📃 statement missing a variable name.",
			"Example: 📃 x 100",
			&symerr.ErrorType{FatalErr: true},
		)
	}
	stmt.Name = p.CurTok.Literal

	p.nextToken()

	// Expect a literal value
	if p.CurTok.Type != tokens.Literal {
		symerr.ErrorMessage(
			"📃 statement missing a literal value.",
			"Example: 📃 <identifier> <literal>",
			&symerr.ErrorType{FatalErr: true},
		)
	}
	// TODO: could convert to int/float if needed
	stmt.Value = p.CurTok.Literal

	p.checkForStop()

	return stmt
}

func (p *Parser) ParseConst() *ConstStatement {
	stmt := &ConstStatement{}

	// Move after declare constant emoji
	p.nextToken()

	// Expect a variable name (identifier)
	if p.CurTok.Type != tokens.Identifier {
		symerr.ErrorMessage(
			"📃 statement missing a variable name.",
			"Example: 📃 x 100",
			&symerr.ErrorType{FatalErr: true},
		)
	}
	stmt.Name = p.CurTok.Literal

	p.nextToken()

	// Expect a literal value
	if p.CurTok.Type != tokens.Literal {
		symerr.ErrorMessage(
			"🪨 statement missing a literal value.",
			"Example: 🪨 <identifier> <literal>",
			&symerr.ErrorType{FatalErr: true},
		)
	}
	// TODO: could convert to int/float if needed
	stmt.Value = p.CurTok.Literal

	p.checkForStop()

	return stmt
}

func (p *Parser) ParseDelete() *DeleteStatement {
	stmt := &DeleteStatement{}

	// move after delete emoji
	p.nextToken()

	// expect an identifier
	if p.CurTok.Type != tokens.Identifier {
		symerr.ErrorMessage(
			"✂️ statement missing a identifier name.",
			"Example: 📃 <identifier>",
			&symerr.ErrorType{FatalErr: true},
		)
	}
	stmt.Name = p.CurTok.Literal

	// then the stop emoji

	p.checkForStop()

	return stmt
}
