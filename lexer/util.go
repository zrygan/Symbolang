package lexer

import (
	"bufio"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/zrygan/symbolang/symerr"
)

func OpenFile(fname string) *Source {
	file, err := os.Open(fname)
	if err != nil {
		symerr.ErrorMessage(
			"Error with opening the file.\n"+err.Error(),
			"",
			&symerr.ErrorType{FatalErr: true},
		)
	}
	defer file.Close()

	var sourceBuilder strings.Builder
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		sourceBuilder.WriteString(line) // append whole line including newline
		if err != nil {
			if err == io.EOF {
				break
			}
			symerr.ErrorMessage(
				"Error reading the file: "+err.Error(),
				"",
				&symerr.ErrorType{FatalErr: true},
			)
		}
	}

	source := sourceBuilder.String()
	return NewSource(source)
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
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
