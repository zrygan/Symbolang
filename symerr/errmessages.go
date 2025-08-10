package symerr

import (
	"fmt"
	"os"
)

type ErrorType struct {
	FatalErr bool
}

func ErrorMessage(message string, hint string, et *ErrorType) {
	fmt.Println("⚠️ Symbolang:", message)
	if hint != "" {
		fmt.Println("\t🔍 Hint:\t", hint)
	}

	if et == nil {
		os.Exit(0)
	} else if et.FatalErr {
		os.Exit(1)
	}
}
