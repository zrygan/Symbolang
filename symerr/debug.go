package symerr

import "fmt"

var globalDebug = true

func DebugMessage(messages ...any) {
	if globalDebug {
		fmt.Println("--- @ ", messages[0])
		for _, message := range messages[1:] {
			fmt.Println("ℹ️\tDebug: ", message)
		}
		fmt.Println("---")
	}
}
