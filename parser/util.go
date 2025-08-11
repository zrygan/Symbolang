package parser

import (
	"github.com/zrygan/symbolang/symerr"
	"github.com/zrygan/symbolang/tokens"
)

func (p *Parser) nextToken() {
	p.CurTok = p.peekTok
	p.peekTok = p.Lexer.ReadNextToken()

	symerr.DebugMessage(
		"nextToken",
		p.CurTok.Type,
		p.CurTok.Literal,
	)
}

func (p *Parser) checkForStop() {
	p.nextToken()

	if p.CurTok.Type != tokens.Stop {
		symerr.ErrorMessage(
			"📃 statement is not closed by a 🫷 emoji.",
			"",
			&symerr.ErrorType{FatalErr: true},
		)
	} else {
		p.nextToken()
	}
}
