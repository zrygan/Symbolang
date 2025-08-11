package parser

var globalEnv = map[string]interface{}{}

func IsIdentifierExist(id string) bool {
	if _, ok := globalEnv[id]; ok {
		return true
	}

	return false
}
