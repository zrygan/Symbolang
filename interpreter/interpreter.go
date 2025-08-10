package interpreter

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/zrygan/symbolang/symerr"
)

func OpenFile(fname string) Source {
	file, err := os.Open(fname)
	if err != nil {
		symerr.ErrorMessage(
			"Error with opening the file. \n"+err.Error(),
			"",
			&symerr.ErrorType{FatalErr: true},
		)
	}

	reader := bufio.NewReader(file)
	var source []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					words := strings.Split(line, " ")
					source = append(source, words...)
				}
				break
			}
			symerr.ErrorMessage(
				"Error reading the file: "+err.Error(),
				"",
				&symerr.ErrorType{FatalErr: true},
			)
		}
		words := strings.Split(line, " ")
		source = append(source, words...)
	}

	return NewSource(source)
}
