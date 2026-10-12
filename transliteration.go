// Package transliteration does translit from Ukrainian to Latin.
//
// According to the rules of transliteration, that are described in the resolution
// of the Cabinet of Ministers of Ukraine №55 dated January 27, 2010.
//
// https://zakon.rada.gov.ua/laws/show/55-2010-%D0%BF#Text
package transliteration

import (
	"strings"
	"unicode"
)

// ukrApostrophe is the Ukrainian apostrophe (U+02BC), which is dropped in transliteration.
const ukrApostrophe = 'ʼ'

// UkrToLat transliterates Ukrainian text into Latin in accordance with established rules
func UkrToLat(ukrainianText string) string {
	characters := []rune(ukrainianText)
	normalizeApostrophes(characters)

	var transliteratedText strings.Builder
	transliteratedText.Grow(len(ukrainianText))

	for i := 0; i < len(characters); i++ {
		currentChar := characters[i]

		latChar, isUkrChar := ukrToLatRules[currentChar]
		if !isUkrChar {
			transliteratedText.WriteRune(currentChar)
			continue
		}

		switch currentChar {
		// Process exceptions for "Зг", "зг"
		case 'З', 'з':
			if i+1 < len(characters) && isZgh(characters[i+1]) {
				transliteratedText.WriteString(zghTransliteration(currentChar, characters[i+1]))
				i++ // the "г" is consumed by the exception
				continue
			}
		// Process exceptions inside word
		case 'є', 'ї', 'Й', 'й', 'ю', 'я':
			if hasUkrainianPrev(characters, i) {
				transliteratedText.WriteString(insideWord[currentChar])
				continue
			}
		// Process uppercase exceptions for vowels
		case 'Ї', 'Є', 'Ю', 'Я':
			if hasUkrainianPrev(characters, i) {
				transliteratedText.WriteString(insideWord[currentChar])
				continue
			}

			if hasUppercaseNeighbour(characters, i) {
				transliteratedText.WriteString(uppercaseVowels[currentChar])
				continue
			}
		// Process uppercase exceptions for consonants
		case 'Ж', 'Х', 'Ц', 'Ч', 'Ш', 'Щ':
			if hasUppercaseNeighbour(characters, i) {
				transliteratedText.WriteString(uppercaseConsonants[currentChar])
				continue
			}
		}

		transliteratedText.WriteString(latChar)
	}

	return transliteratedText.String()
}

// normalizeApostrophes replaces ASCII apostrophe and typographic apostrophe with the
// Ukrainian apostrophe ʼ when it follows a Ukrainian letter, so that it is dropped
// the same way and does not hide the preceding letter from the "inside word" rules.
func normalizeApostrophes(characters []rune) {
	for i := 1; i < len(characters); i++ {
		if isApostrophe(characters[i]) && isUkrainian(characters[i-1]) {
			characters[i] = ukrApostrophe
		}
	}
}

func isApostrophe(r rune) bool {
	return r == '\'' || r == '’' || r == ukrApostrophe
}

func isUkrainian(r rune) bool {
	_, ok := ukrToLatRules[r]
	return ok
}

// hasUkrainianPrev reports whether the character before position i is a Ukrainian letter.
func hasUkrainianPrev(characters []rune, i int) bool {
	return i > 0 && isUkrainian(characters[i-1])
}

// hasUppercaseNeighbour reports whether the character before or after position i is uppercase.
func hasUppercaseNeighbour(characters []rune, i int) bool {
	return (i > 0 && unicode.IsUpper(characters[i-1])) ||
		(i+1 < len(characters) && unicode.IsUpper(characters[i+1]))
}

func isZgh(r rune) bool {
	return r == 'Г' || r == 'г'
}

// zghTransliteration returns "Zgh", "zgh", "ZGH" or "zGH" depending on the case of both letters.
func zghTransliteration(first, second rune) string {
	result := "zgh"
	if unicode.IsUpper(first) {
		result = "Z" + result[1:]
	}
	if unicode.IsUpper(second) {
		result = result[:1] + "GH"
	}
	return result
}
