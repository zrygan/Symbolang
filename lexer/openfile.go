package lexer

import (
	"bufio"
	"io"
	"os"
	"strings"

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
