package symerr

import (
	"fmt"
	"os"
)

type ErrorType struct {
	FatalErr bool
	ConstErr bool
	DelErr   bool
}

func ErrorMessage(message string, hint string, et *ErrorType) {
	fmt.Println("⚠️ Symbolang:", message)
	if hint != "" {
		fmt.Println("\t🔍 Hint: ", hint)
	}

	if et == nil {
		os.Exit(0)
	} else if et.FatalErr {
		os.Exit(1)
	} else if et.ConstErr {
		return
	}
}
