package tokens

type TokenType int

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
	// General tokens
	Literal TokenType = iota
	Identifier
	Illegal
	EOF

	// Stop token
	Stop

	// Input-Output tokens
	Pencil
	Pen

	// Variable and Constant tokens
	Paper
	Rock
	Scissor
)

const (
	// General Emojis
	EmojiStop = "🫷"

	// Input-Output Emojis
	EmojiPencil = "✏️"
	EmojiPen    = "🖊️"

	// Variable and Constant Emojis
	EmojiPaper   = "📃"
	EmojiRock    = "🪨"
	EmojiScissor = "✂️"
)

var EMOJI_TOKENS = []string{
	EmojiStop,
	EmojiPencil,
	EmojiPen,
	EmojiPaper,
	EmojiRock,
	EmojiScissor,
}
