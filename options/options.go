package options

import (
	"github.com/zrygan/symbolang/symerr"
)

var interpreterOptions = map[string]func(){
	"-dp": func() {
		symerr.DebugFlags["ParseProgram"] = true
	},
	"-dt": func() {
		symerr.DebugFlags["NextToken"] = true
	},
	"-dall": func() {
		for flag := range symerr.DebugFlags {
			symerr.DebugFlags[flag] = true
		}
	},
}

func Options(opts []string) {
	for _, opt := range opts {
		f, ok := interpreterOptions[opt]

		if !ok {
			symerr.ErrorMessage(
				"Interpreter option"+opt+"is not a valid option",
				"Use `symbolang ??` to get a list of options",
				nil,
			)
		}

		f()
	}
}
