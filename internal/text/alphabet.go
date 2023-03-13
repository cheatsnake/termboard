package text

import "strings"

const (
	englishLetters = "abcdefghijklmnopqrstuvwxyz"
	russianLetters = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	numbers        = "0123456789"
	symbols        = "!@#$%&?-+=~(){}[]., "
)

// Alphabet return a slice of charactes for given language
func Alphabet(lang string) []rune {
	switch lang {
	case "ru":
		return generate(russianLetters)
	default:
		return generate(englishLetters)
	}
}

func generate(letters string) []rune {
	chars := strings.Join(
		[]string{
			letters,
			strings.ToUpper(letters),
			numbers,
			symbols,
		}, "",
	)
	alphabet := make([]rune, 0, len(chars))

	for _, char := range chars {
		alphabet = append(alphabet, char)
	}

	return alphabet
}
