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

func (p *Parser) nextToken() {
	p.CurTok = p.peekTok
	p.peekTok = p.Lexer.ReadNextToken()
}

func (p *Parser) ParseProgram() []Statement {
	statements := []Statement{}

	for p.CurTok.Type != tokens.EOF {
		switch p.CurTok.Type {
		case tokens.Pencil, tokens.Pen:
			stmt := p.ParsePrintStatement()
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
	p.nextToken()

	if p.CurTok.Type != tokens.Literal {
		symerr.ErrorMessage(
			"A ✏️ or 🖊️ statement does not have a literal input.",
			"",
			&symerr.ErrorType{FatalErr: true},
		)
	}
	stmt.Value = p.CurTok.Literal

	p.nextToken()
	if p.CurTok.Type != tokens.Stop {
		symerr.ErrorMessage(
			"A ✏️ or 🖊️ statement is not closed by a 🫷 emoji.",
			"",
			&symerr.ErrorType{FatalErr: true},
		)
	} else {
		p.nextToken()
	}

	return stmt
}
