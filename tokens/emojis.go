package tokens

import "strings"

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
			case EmojiPaper:
				return emoji, Paper, true
			case EmojiRock:
				return emoji, Rock, true
			case EmojiScissor:
				return emoji, Scissor, true
			}
		}
	}

	return "", Literal, false
}
