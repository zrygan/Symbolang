package interpreter

import "strings"

type TokenType int

const (
	// General tokens
	Literal TokenType = iota
	EOF
	Stop

	// Input-Output tokens
	Pencil
	Pen
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokType TokenType, literal string) Token {
	return Token{
		tokType,
		literal,
	}
}

const (
	// General Emojis
	EmojiStop = "🫷"

	// Input-Output Emojis
	EmojiPencil = "✏️"
	EmojiPen    = "🖊️"
)

var EMOJI_TOKENS = []string{
	EmojiStop,
	EmojiPencil,
	EmojiPen,
}

// IsEmojiAt() gets the emoji (the symbol and token) if it exists.
func IsEmojiAt(input string, pos int) (string, TokenType, bool) {
	for _, emoji := range EMOJI_TOKENS {
		if strings.HasPrefix(input[pos:], emoji) {
			switch emoji {
			case EmojiStop:
				return emoji, Stop, true
			case EmojiPencil:
				return emoji, Pencil, true
			case EmojiPen:
				return emoji, Pen, true
			}
		}
	}

	return "", Literal, false
}
