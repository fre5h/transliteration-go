/*
	Package transliteration provides methods for transliterating text from Ukrainian to Latin.

	According to the rules of Ukrainian-to-Latin transliteration, that are described in the resolution
	of the Cabinet of Ministers of Ukraine №55 dated January 27, 2010.

	https://zakon.rada.gov.ua/laws/show/55-2010-%D0%BF#Text
*/
package transliteration

var ukrToLatRules = map[rune]string{
	'А': "A",
	'Б': "B",
	'В': "V",
	'Г': "H",
	'Ґ': "G",
	'Д': "D",
	'Е': "E",
	'Є': "Ye",
	'Ж': "Zh",
	'З': "Z",
	'И': "Y",
	'І': "I",
	'Ї': "Yi",
	'Й': "Y",
	'К': "K",
	'Л': "L",
	'М': "M",
	'Н': "N",
	'О': "O",
	'П': "P",
	'Р': "R",
	'С': "S",
	'Т': "T",
	'У': "U",
	'Ф': "F",
	'Х': "Kh",
	'Ц': "Ts",
	'Ч': "Ch",
	'Ш': "Sh",
	'Щ': "Shch",
	'Ь': "",
	'Ю': "Yu",
	'Я': "Ya",
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "h",
	'ґ': "g",
	'д': "d",
	'е': "e",
	'є': "ye",
	'ж': "zh",
	'з': "z",
	'и': "y",
	'і': "i",
	'ї': "yi",
	'й': "y",
	'к': "k",
	'л': "l",
	'м': "m",
	'н': "n",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'у': "u",
	'ф': "f",
	'х': "kh",
	'ц': "ts",
	'ч': "ch",
	'ш': "sh",
	'щ': "shch",
	'ь': "",
	'ю': "yu",
	'я': "ya",
	'ʼ': "",
}

var insideWord = map[rune]string{
	'Ї': "I",
	'Й': "I",
	'Є': "IE",
	'Ю': "IU",
	'Я': "IA",
	'ї': "i",
	'й': "i",
	'є': "ie",
	'ю': "iu",
	'я': "ia",
}

var uppercaseVowels = map[rune]string{
	'Ї': "YI",
	'Є': "YE",
	'Ю': "YU",
	'Я': "YA",
}

var uppercaseConsonants = map[rune]string{
	'Ж': "ZH",
	'Х': "KH",
	'Ц': "TS",
	'Ч': "CH",
	'Ш': "SH",
	'Щ': "SHCH",
}
