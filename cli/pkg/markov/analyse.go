package markov

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// AnalyseLine analyses the line text and updates the model
func (m *Model) AnalyseLine(line string) error {
	if strings.ContainsRune(line, EOLValue) == true {
		return errors.New("line cannot contain newlines")
	}

	words := strings.Split(line, " ")
	for i, word := range words {
		var next string
		if i == (len(words) - 1) {
			next = string(EOLValue)
		} else {
			next = words[i+1]
		}

		token, ok := m.Tokens[word]
		if ok == false {
			token = NewToken(word)
		}

		token.Next.Add(next)
		if i == 0 {
			m.StartTokens.Add(word)
		}

		m.Tokens[word] = token
	}

	return nil
}

// AnalyseMany analyses a block of text, splitting it on every newline
// Each line is analysed as an individual unit
func (m *Model) AnalyseMany(text string) error {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		err := m.AnalyseLine(line)

		if err != nil {
			return err
		}
	}

	return nil
}

// AnalyseReader analyses each line of an io.Reader
func (m *Model) AnalyseReader(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		err := m.AnalyseLine(scanner.Text())

		if err != nil {
			return err
		}
	}

	return nil
}
