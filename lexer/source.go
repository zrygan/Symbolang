package lexer

type Source struct {
	SourceFile string
}

func NewSource(code string) *Source {
	return &Source{
		code,
	}
}

func (s *Source) DebugSource() {
	println(s.SourceFile)
}
