package markov

type wordMap map[string]int

// Add adds the word to the map
// Each successive call with the same word will increment the weight
func (w wordMap) Add(word string) {
	weight := w[word]
	w[word] = weight + 1
}

// Set sets the weight of a word in the map
func (w wordMap) Set(word string, weight int) {
	w[word] = weight
}

// Delete removes a word from the map
func (w wordMap) Delete(word string) {
	delete(w, word)
}

// Select selects a token at random, but taking the weight of all words into account
func (w wordMap) Select() string {
	keys := []string{}
	weights := []int{}

	for key, weight := range w {
		keys = append(keys, key)
		weights = append(weights, weight)
	}

	idx := weightedRandomIdx(weights)
	return keys[idx]
}

// All returns all words currently stored
func (w wordMap) All() []string {
	keys := make([]string, len(w))

	i := 0
	for k := range w {
		keys[i] = k
		i++
	}

	return keys
}

// Get returns the weight and existence of a specific word
func (w wordMap) Get(word string) (int, bool) {
	weight, ok := w[word]
	return weight, ok
}
