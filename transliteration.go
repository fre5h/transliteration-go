// Transliteration from Ukrainian to Latin.
//
// According to the rules of transliteration, that are described in the resolution
// of the Cabinet of Ministers of Ukraine №55 dated January 27, 2010.
//
// https://zakon.rada.gov.ua/laws/show/55-2010-%D0%BF#Text

package transliteration

import "unicode"

func UkrToLat(ukrainianText string) (transliteratedText string) {
	characters := []rune(ukrainianText)
	numberOfChars := len(characters)
	skipNextChar := false

	for i, currentChar := range characters {
		if skipNextChar {
			skipNextChar = false
			continue
		}

		if latChar, isUrkChar := ukrToLatRules[currentChar]; isUrkChar {
			switch currentChar {
			// Process exceptions for "Зг", "зг"
			case 'З', 'з':
				if i+1 != numberOfChars { // has next char
					nextChar := characters[i+1]
					if nextChar == 'Г' || nextChar == 'г' {
						transliteratedText += processZghException(currentChar, nextChar)
						skipNextChar = true
						continue
					}
				}
			// Process exceptions inside word
			case 'є', 'ї', 'Й', 'й', 'ю', 'я':
				// Has prev char
				if i != 0 {
					if _, prevCharacterIsUkrainian := ukrToLatRules[characters[i-1]]; prevCharacterIsUkrainian {
						transliteratedText += insideWord[currentChar]
						continue
					}
				}
			// Process uppercase exceptions for vowels
			case 'Ї', 'Є', 'Ю', 'Я':
				// Has prev char
				if i != 0 {
					if _, prevCharIsUkr := ukrToLatRules[characters[i-1]]; prevCharIsUkr {
						transliteratedText += insideWord[currentChar]
						continue
					}
				}

				// Previous or next character is also uppercase
				if (i != 0 && unicode.IsUpper(characters[i-1])) || (i+1 != numberOfChars && unicode.IsUpper(characters[i+1])) {
					transliteratedText += uppercaseVowels[currentChar]
					continue
				}
			// Process uppercase exceptions for consonants
			case 'Ж', 'Х', 'Ц', 'Ч', 'Ш', 'Щ':
				// Previous or next character is also uppercase
				if (i != 0 && unicode.IsUpper(characters[i-1])) || (i+1 != numberOfChars && unicode.IsUpper(characters[i+1])) {
					transliteratedText += uppercaseConsonants[currentChar]
					continue
				}
			}

			transliteratedText += latChar
		} else {
			transliteratedText += string(currentChar)
		}
	}

	return
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
