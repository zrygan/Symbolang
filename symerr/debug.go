package symerr

import "fmt"

// Debugging Booleans
var DebugFlags = map[string]bool{
	"ParseProgram": false,
	"NextToken":    false,
}

func DebugMessage(messages ...any) {
	fmt.Println("--- @ ", messages[0])
	for _, message := range messages[1:] {
		fmt.Println("ℹ️\tDebug: ", message)
	}
	fmt.Println("---")
}
