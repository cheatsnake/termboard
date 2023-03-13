package text

import (
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed data/words.txt
var wordsFile string
var words = loadWords()

// RandomText generate random string with given number of words
func RandomText(amount int) string {
	randWords := randomWords(words, amount)
	return strings.Join(randWords, " ")
}

func randomWords(w []string, amount int) []string {
	words := make([]string, amount)

	for i := 0; i < amount; i++ {
		words[i] = w[rand.Intn(len(w))]
	}

	return words
}

func loadWords() []string {
	return strings.Split(wordsFile, "\n")
}
