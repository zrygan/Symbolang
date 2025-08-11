package parser

import "github.com/zrygan/symbolang/symerr"

func (p *Parser) nextToken() {
	p.CurTok = p.peekTok
	p.peekTok = p.Lexer.ReadNextToken()

	symerr.DebugMessage(
		"nextToken",
		p.CurTok.Type,
		p.CurTok.Literal,
	)
}
