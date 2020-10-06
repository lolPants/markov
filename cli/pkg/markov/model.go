package markov

import (
	"encoding/json"
	"io"
	"strings"
)

// NewModel creates a new Markov Model
func NewModel() *Model {
	return &Model{
		Tokens:      make(map[string]*Token),
		StartTokens: make(wordMap),
	}
}

// Model represents a markov chain model
type Model struct {
	Tokens      map[string]*Token `json:"tokens"`
	StartTokens wordMap           `json:"startTokens"`
}

// Generate a string using the markov model
func (m *Model) Generate() string {
	start := m.StartTokens.Select()
	token := m.Tokens[start]

	var sb strings.Builder
	first := true

	for {
		if token == nil {
			break
		}

		if first == true {
			first = false
		} else {
			sb.WriteRune(' ')
		}

		sb.WriteString(token.Value)
		token = token.SelectNext(m)
	}

	return sb.String()
}

// Export exports the model as a JSON byte slice
func (m *Model) Export() ([]byte, error) {
	bytes, err := json.Marshal(m)
	return bytes, err
}

// ExportWriter exports the model JSON to the passed writer
func (m *Model) ExportWriter(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(m)
}

// Import imports a model from a JSON byte slice
//
// WARNING: This will overwrite the current model
func (m *Model) Import(data []byte) error {
	return json.Unmarshal(data, m)
}

// ImportReader imports a model from an io.Reader to a JSON document
//
// WARNING: This will overwrite the current model
func (m *Model) ImportReader(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(m)
}
