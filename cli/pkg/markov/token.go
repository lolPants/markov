package markov

const (
	// EOLValue is used to denote the end of a chain
	EOLValue = '\n'
)

// NewToken creates a new token
func NewToken(value string) *Token {
	return &Token{
		Value: value,
		Next:  make(wordMap),
	}
}

// Token represents a word token
type Token struct {
	Value string  `json:"value"`
	Next  wordMap `json:"next"`
}

// SelectNext selects the next token according to weighted randomness
func (t *Token) SelectNext(m *Model) *Token {
	value := t.Next.Select()

	token := m.Tokens[value]
	return token
}
