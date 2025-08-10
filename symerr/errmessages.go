package symerr

import (
	"fmt"
	"os"
)

type ErrorType struct {
	ReadError bool
}

func ErrorMessage(message string, hint string, et *ErrorType) {
	fmt.Println("⚠️ Symbolang:", message)
	if hint != "" {
		fmt.Println("\t🔍 Hint:\t", hint)
	}

	if et.ReadError {
		os.Exit(1)
	}
}
