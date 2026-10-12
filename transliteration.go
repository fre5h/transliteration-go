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
	numberOfChars := len(characters)
	skipNextChar := false

	var transliteratedText strings.Builder
	transliteratedText.Grow(len(ukrainianText))

	for i, currentChar := range characters {
		if skipNextChar {
			skipNextChar = false
			continue
		}

		if latChar, isUkrChar := ukrToLatRules[currentChar]; isUkrChar {
			switch currentChar {
			// Process exceptions for "Зг", "зг"
			case 'З', 'з':
				if i+1 != numberOfChars { // has next char
					nextChar := characters[i+1]
					if nextChar == 'Г' || nextChar == 'г' {
						transliteratedText.WriteString(processZghException(currentChar, nextChar))
						skipNextChar = true
						continue
					}
				}
			// Process exceptions inside word
			case 'є', 'ї', 'Й', 'й', 'ю', 'я':
				// Has prev char
				if i != 0 {
					if _, prevCharacterIsUkrainian := ukrToLatRules[characters[i-1]]; prevCharacterIsUkrainian {
						transliteratedText.WriteString(insideWord[currentChar])
						continue
					}
				}
			// Process uppercase exceptions for vowels
			case 'Ї', 'Є', 'Ю', 'Я':
				// Has prev char
				if i != 0 {
					if _, prevCharIsUkr := ukrToLatRules[characters[i-1]]; prevCharIsUkr {
						transliteratedText.WriteString(insideWord[currentChar])
						continue
					}
				}

				// Previous or next character is also uppercase
				if (i != 0 && unicode.IsUpper(characters[i-1])) || (i+1 != numberOfChars && unicode.IsUpper(characters[i+1])) {
					transliteratedText.WriteString(uppercaseVowels[currentChar])
					continue
				}
			// Process uppercase exceptions for consonants
			case 'Ж', 'Х', 'Ц', 'Ч', 'Ш', 'Щ':
				// Previous or next character is also uppercase
				if (i != 0 && unicode.IsUpper(characters[i-1])) || (i+1 != numberOfChars && unicode.IsUpper(characters[i+1])) {
					transliteratedText.WriteString(uppercaseConsonants[currentChar])
					continue
				}
			}

			transliteratedText.WriteString(latChar)
		} else {
			transliteratedText.WriteRune(currentChar)
		}
	}

	return transliteratedText.String()
}

// normalizeApostrophes replaces ASCII apostrophe and typographic apostrophe with the
// Ukrainian apostrophe ʼ when it follows a Ukrainian letter, so that it is dropped
// the same way and does not hide the preceding letter from the "inside word" rules.
func normalizeApostrophes(characters []rune) {
	for i := 1; i < len(characters); i++ {
		if !isApostrophe(characters[i]) {
			continue
		}
		if _, prevIsUkr := ukrToLatRules[characters[i-1]]; prevIsUkr {
			characters[i] = ukrApostrophe
		}
	}
}

func isApostrophe(r rune) bool {
	return r == '\'' || r == '’' || r == ukrApostrophe
}

func processZghException(firstChar, secondChar rune) (result string) {
	if firstChar == 'З' && secondChar == 'г' {
		result = "Zgh"
	} else if firstChar == 'З' && secondChar == 'Г' {
		result = "ZGH"
	} else if firstChar == 'з' && secondChar == 'г' {
		result = "zgh"
	} else if firstChar == 'з' && secondChar == 'Г' {
		result = "zGH"
	}

	return
}
