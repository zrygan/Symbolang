package interpreter

import (
	"unicode"
	"unicode/utf8"
)

type Lex struct {
	input   string
	pos     int // current byte position in input
	readPos int // next reading position (byte index)
	c       rune
}

func NewLex(input string) *Lex {
	l := &Lex{input: input}
	l.readChar()
	return l
}

// readChar() reads the current character at the lexer pointer then moves it.
func (l *Lex) readChar() {
	if l.readPos >= len(l.input) {
		l.c = 0
	} else {
		r, width := utf8.DecodeRuneInString(l.input[l.readPos:])
		l.c = r
		l.pos = l.readPos
		l.readPos += width
	}
}

// readStringLiteral() captures the entire literal enclosed by " ". Then, it
// moves the lexer pointer to the character after the closing ".
func (l *Lex) readStringLiteral() string {
	start := l.pos
	l.readChar()

	for l.c != 0 && l.c != '"' {
		l.readChar()
	}

	// if the closing " is found, end here
	if l.c == '"' {
		l.readChar()
	}

	return l.input[start:l.pos]
}

// ReadNextToken() uses the lexer to read the next token.
// Tokens are available at interpreter/tokens.go
func (l *Lex) ReadNextToken() Token {
	for unicode.IsSpace(l.c) {
		l.readChar()
	}

	if l.c == 0 {
		return NewToken(EOF, "")
	}

	// check if this token is an emoji (we handle this with care since emojis
	// are NOT runes or single width characters)!
	emoji, tokType, isEmoji := IsEmojiAt(l.input, l.pos)
	if isEmoji {
		l.advanceBy(len(emoji))
		return NewToken(tokType, emoji)
	}

	// handle string literals
	if l.c == '"' {
		str := l.readStringLiteral()
		return NewToken(Literal, str)
	}

	// for other literals (numbers and variables)
	start := l.pos
	for l.c != 0 && !unicode.IsSpace(l.c) && !isEmoji {
		l.readChar()
	}
	literal := l.input[start:l.pos]
	return NewToken(Literal, literal)
}

// advanceBy() moved the lexer pointer by n units or characters.
func (l *Lex) advanceBy(n int) {
	l.pos += n
	l.readPos = l.pos
	if l.pos >= len(l.input) {
		l.c = 0
		return
	}
	r, _ := utf8.DecodeRuneInString(l.input[l.pos:])
	l.c = r
}
