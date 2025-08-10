package interpreter

import "strconv"

type Source struct {
	SourceFile []string
}

func NewSource(code []string) Source {
	return Source{
		code,
	}
}

func (s *Source) DebugSource() {
	for num, word := range s.SourceFile {
		println(strconv.Itoa(num) + "\t" + word)
	}
}
