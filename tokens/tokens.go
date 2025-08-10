package tokens

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
