// Package transliteration contains toolkit for translating strings from Ukrainian Cyrillic to Ukrainian Latin.
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

// CyrToLat transliterates Cyrillic line of text into Latin in accordance with established rules
func CyrToLat(cyrillicLine string) string {
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

// CyrWordToLat transliterates Cyrillic word into Latin in accordance with established rules
func CyrWordToLat(cyrillicWord string) string {
	var latinWord string
	skipNextRune := false
	cyrillicRunes := []rune(cyrillicWord)

	for indexOfRune, cyrillicRune := range cyrillicRunes {
		if skipNextRune {
			skipNextRune = false
			continue
		}
		var latinStr string

		// these pointers allow to recognise presence of a rune without checking the slice each time
		var nextRune, previousRune *rune

		if indexOfRune+1 != len(cyrillicRunes) {
			nextRune = &cyrillicRunes[indexOfRune+1]
		}

		if indexOfRune != 0 {
			previousRune = &cyrillicRunes[indexOfRune-1]
		}

		latinStr, skipNextRune = CyrRuneToLatString(cyrillicRune, nextRune, previousRune)
		latinWord += latinStr
	}

	return latinWord
}

// CyrRuneToLatString transliterates Cyrillic rune into Latin string (sequence of runes). Additionally, it returns bool of skip for the text rune
func CyrRuneToLatString(cyrillicRune rune, nextRune, previousRune *rune) (string, bool) {
	switch cyrillicRune {
	case 'З', 'з': // Process exceptions for "Зг", "зг"
		if nextRune == nil {
			break
		}

		if *nextRune == 'Г' || *nextRune == 'г' {
			return zghEdgecases[[2]rune{cyrillicRune, *nextRune}], true
		}

	case 'є', 'ї', 'Й', 'й', 'ю', 'я': // Process exceptions inside word
		if previousRune == nil {
			break
		}

		_, previousIsCyrillic := cyrillicToLatin[*previousRune]
		if previousIsCyrillic {
			return insideWord[cyrillicRune], false
		}

	case 'Ї', 'Є', 'Ю', 'Я': // Process uppercase exceptions for vowels
		if previousRune != nil {
			_, previousIsCyrillic := cyrillicToLatin[*previousRune]
			if previousIsCyrillic {
				return insideWord[cyrillicRune], false
			}
		}
		// Previous or next character is also uppercase
		if (previousRune != nil && unicode.IsUpper(*previousRune)) || (nextRune != nil && unicode.IsUpper(*nextRune)) {
			return uppercaseVowels[cyrillicRune], false
		}

	case 'Ж', 'Х', 'Ц', 'Ч', 'Ш', 'Щ': // Process uppercase exceptions for consonants
		// Previous or next character is also uppercase
		if (previousRune != nil && unicode.IsUpper(*previousRune)) || (nextRune != nil && unicode.IsUpper(*nextRune)) {
			return uppercaseConsonants[cyrillicRune], false
		}
	}

	latinMapped, hasLatin := cyrillicToLatin[cyrillicRune]
	if !hasLatin {
		return string(cyrillicRune), false
	}
	return latinMapped, false
}
