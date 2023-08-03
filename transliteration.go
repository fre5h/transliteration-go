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

// CyrToLat transliterates Ukrainian text into Latin in accordance with established rules
func CyrToLat(cyrillicText string) string {
	var latinText string

	cyrillicLines := strings.Split(cyrillicText, "\n")

	for _, cyrillicLine := range cyrillicLines {
		latinLine := CyrLineToLat(cyrillicLine)
		latinText += latinLine
	}

	return latinText
}

func CyrLineToLat(cyrillicLine string) string {
	var latinLine string

	cyrillicWords := strings.Split(cyrillicLine, " ")

	for i, cyrillicWord := range cyrillicWords {
		latinWord := CyrWordToLat(cyrillicWord)
		latinLine += latinWord

		if i != len(cyrillicWords)-1 {
			latinLine += " "
		}
	}

	return latinLine
}

func CyrWordToLat(cyrillicWord string) string {
	var latinWord []rune
	skipNextChar := false
	cyrillicRunes := []rune(cyrillicWord)

	for i, cyrillicRune := range cyrillicRunes {
		if skipNextChar {
			skipNextChar = false
			continue
		}

		latinMapped, hasLatin := cyrillicToLatin[cyrillicRune]
		if !hasLatin {
			latinWord = append(latinWord, cyrillicRune)
		}

		switch cyrillicRune {
		// Process exceptions for "Зг", "зг"
		case 'З', 'з':
			if i+1 == len(cyrillicRunes) { // has next char
				break
			}

			nextRune := cyrillicRunes[i+1]
			if nextRune == 'Г' || nextRune == 'г' {
				latinWord = append(latinWord, []rune(zghEdgecases[[2]rune{cyrillicRune, nextRune}])...)
				skipNextChar = true
				continue
			}
		// Process exceptions inside word
		case 'є', 'ї', 'Й', 'й', 'ю', 'я':
			// Has prev char
			if i != 0 {
				if _, prevCharacterIsUkrainian := cyrillicToLatin[cyrillicRunes[i-1]]; prevCharacterIsUkrainian {
					latinWord = append(latinWord, []rune(insideWord[cyrillicRune])...)
					continue
				}
			}
		// Process uppercase exceptions for vowels
		case 'Ї', 'Є', 'Ю', 'Я':
			// Has prev char
			if i != 0 {
				if _, prevCharIsUkr := cyrillicToLatin[cyrillicRunes[i-1]]; prevCharIsUkr {
					latinWord = append(latinWord, []rune(insideWord[cyrillicRune])...)
					continue
				}
			}

			// Previous or next character is also uppercase
			if (i != 0 && unicode.IsUpper(cyrillicRunes[i-1])) || (i+1 != len(cyrillicRunes) && unicode.IsUpper(cyrillicRunes[i+1])) {
				latinWord = append(latinWord, []rune(uppercaseVowels[cyrillicRune])...)
				continue
			}
		// Process uppercase exceptions for consonants
		case 'Ж', 'Х', 'Ц', 'Ч', 'Ш', 'Щ':
			// Previous or next character is also uppercase
			if (i != 0 && unicode.IsUpper(cyrillicRunes[i-1])) || (i+1 != len(cyrillicRunes) && unicode.IsUpper(cyrillicRunes[i+1])) {
				latinWord = append(latinWord, []rune(uppercaseConsonants[cyrillicRune])...)
				continue
			}
		}

		latinWord = append(latinWord, []rune(latinMapped)...)
	}

	return string(latinWord)
}
