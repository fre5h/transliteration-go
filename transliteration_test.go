package transliteration_test

import (
	"fmt"
	"testing"

	"github.com/fre5h/transliteration-go"
)

const reset = "\033[0m"
const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"

func TestUkrToLat(t *testing.T) {
	for _, data := range testVariants {
		if result := transliteration.UkrToLat(data.ukrainian); result != data.latin {
			t.Errorf(
				"Transliteration of %s\"%s\"%s is incorrect, expected: %s\"%s\"%s, actual: %s\"%s\"%s.",
				yellow, data.ukrainian, reset,
				green, data.latin, reset,
				red, result, reset,
			)
		}
	}
}

func ExampleUkrToLat() {
	fmt.Println(transliteration.UkrToLat("Слава Україні!"))
	// Output:
	// Slava Ukraini!
}

var testVariants = []struct {
	ukrainian string
	latin     string
}{
	// Ukrainian Alphabet
	{"А", "A"},
	{"Б", "B"},
	{"В", "V"},
	{"Г", "H"},
	{"Ґ", "G"},
	{"Д", "D"},
	{"Е", "E"},
	{"Є", "Ye"},
	{"Ж", "Zh"},
	{"З", "Z"},
	{"И", "Y"},
	{"І", "I"},
	{"Ї", "Yi"},
	{"Й", "Y"},
	{"К", "K"},
	{"Л", "L"},
	{"М", "M"},
	{"Н", "N"},
	{"О", "O"},
	{"П", "P"},
	{"Р", "R"},
	{"С", "S"},
	{"Т", "T"},
	{"У", "U"},
	{"Ф", "F"},
	{"Х", "Kh"},
	{"Ц", "Ts"},
	{"Ч", "Ch"},
	{"Ш", "Sh"},
	{"Щ", "Shch"},
	{"Ю", "Yu"},
	{"Я", "Ya"},
	{"а", "a"},
	{"б", "b"},
	{"в", "v"},
	{"г", "h"},
	{"ґ", "g"},
	{"д", "d"},
	{"е", "e"},
	{"є", "ye"},
	{"ж", "zh"},
	{"з", "z"},
	{"и", "y"},
	{"і", "i"},
	{"ї", "yi"},
	{"й", "y"},
	{"к", "k"},
	{"л", "l"},
	{"м", "m"},
	{"н", "n"},
	{"о", "o"},
	{"п", "p"},
	{"р", "r"},
	{"с", "s"},
	{"т", "t"},
	{"у", "u"},
	{"ф", "f"},
	{"х", "kh"},
	{"ц", "ts"},
	{"ч", "ch"},
	{"ш", "sh"},
	{"щ", "shch"},
	{"ь", ""},
	{"ю", "yu"},
	{"я", "ya"},
	{"ʼ", ""},

	// Examples of transliteration form the resolution of the Cabinet of Ministers of Ukraine №55 (27.01.2010)
	// Аа
	{"Алушта", "Alushta"},
	{"АЛУШТА", "ALUSHTA"},
	{"алушта", "alushta"},
	{"Андрій", "Andrii"},
	{"АНДРІЙ", "ANDRII"},
	{"андрій", "andrii"},
	// Бб
	{"Борщагівка", "Borshchahivka"},
	{"БОРЩАГІВКА", "BORSHCHAHIVKA"},
	{"борщагівка", "borshchahivka"},
	{"Борисенко", "Borysenko"},
	{"БОРИСЕНКО", "BORYSENKO"},
	{"борисенко", "borysenko"},
	// Вв
	{"Вінниця", "Vinnytsia"},
	{"ВІННИЦЯ", "VINNYTSIA"},
	{"вінниця", "vinnytsia"},
	{"Володимир", "Volodymyr"},
	{"ВОЛОДИМИР", "VOLODYMYR"},
	{"володимир", "volodymyr"},
	// Гг
	{"Гадяч", "Hadiach"},
	{"ГАДЯЧ", "HADIACH"},
	{"гадяч", "hadiach"},
	{"Богдан", "Bohdan"},
	{"БОГДАН", "BOHDAN"},
	{"богдан", "bohdan"},
	{"Згурський", "Zghurskyi"},
	{"ЗГУРСЬКИЙ", "ZGHURSKYI"},
	{"згурський", "zghurskyi"},
	// Ґґ
	{"Ґалаґан", "Galagan"},
	{"ҐАЛАҐАН", "GALAGAN"},
	{"ґалаґан", "galagan"},
	{"Ґорґани", "Gorgany"},
	{"ҐОРҐАНИ", "GORGANY"},
	{"ґорґани", "gorgany"},
	// Дд
	{"Донецьк", "Donetsk"},
	{"ДОНЕЦЬК", "DONETSK"},
	{"донецьк", "donetsk"},
	{"Дмитро", "Dmytro"},
	{"ДМИТРО", "DMYTRO"},
	{"дмитро", "dmytro"},
	// Ее
	{"Рівне", "Rivne"},
	{"РІВНЕ", "RIVNE"},
	{"рівне", "rivne"},
	{"Олег", "Oleh"},
	{"ОЛЕГ", "OLEH"},
	{"олег", "oleh"},
	{"Есмань", "Esman"},
	{"ЕСМАНЬ", "ESMAN"},
	{"есмань", "esman"},
	// Єє
	{"Єнакієве", "Yenakiieve"},
	{"ЄНАКІЄВЕ", "YENAKIIEVE"},
	{"єнакієве", "yenakiieve"},
	{"Гаєвич", "Haievych"},
	{"ГАЄВИЧ", "HAIEVYCH"},
	{"гаєвич", "haievych"},
	{"Коропʼє", "Koropie"},
	{"КОРОПʼЄ", "KOROPIE"},
	{"коропʼє", "koropie"},
	// Жж
	{"Житомир", "Zhytomyr"},
	{"ЖИТОМИР", "ZHYTOMYR"},
	{"житомир", "zhytomyr"},
	{"Жанна", "Zhanna"},
	{"ЖАННА", "ZHANNA"},
	{"жанна", "zhanna"},
	{"Жежелів", "Zhezheliv"},
	{"ЖЕЖЕЛІВ", "ZHEZHELIV"},
	{"жежелів", "zhezheliv"},
	// Зз
	{"Закарпаття", "Zakarpattia"},
	{"ЗАКАРПАТТЯ", "ZAKARPATTIA"},
	{"закарпаття", "zakarpattia"},
	{"Казимирчук", "Kazymyrchuk"},
	{"КАЗИМИРЧУК", "KAZYMYRCHUK"},
	{"казимирчук", "kazymyrchuk"},
	// Ии
	{"Медвин", "Medvyn"},
	{"МЕДВИН", "MEDVYN"},
	{"медвин", "medvyn"},
	{"Михайленко", "Mykhailenko"},
	{"МИХАЙЛЕНКО", "MYKHAILENKO"},
	{"михайленко", "mykhailenko"},
	// Іі
	{"Іванків", "Ivankiv"},
	{"ІВАНКІВ", "IVANKIV"},
	{"іванків", "ivankiv"},
	{"Іващенко", "Ivashchenko"},
	{"ІВАЩЕНКО", "IVASHCHENKO"},
	{"іващенко", "ivashchenko"},
	// Її
	{"Їжакевич", "Yizhakevych"},
	{"ЇЖАКЕВИЧ", "YIZHAKEVYCH"},
	{"їжакевич", "yizhakevych"},
	{"Кадиївка", "Kadyivka"},
	{"КАДИЇВКА", "KADYIVKA"},
	{"кадиївка", "kadyivka"},
	{"Марʼїне", "Marine"},
	{"МАРʼЇНЕ", "MARINE"},
	{"марʼїне", "marine"},
	// Йй
	{"Йосипівка", "Yosypivka"},
	{"ЙОСИПІВКА", "YOSYPIVKA"},
	{"йосипівка", "yosypivka"},
	{"Стрий", "Stryi"},
	{"СТРИЙ", "STRYI"},
	{"стрий", "stryi"},
	{"Олексій", "Oleksii"},
	{"ОЛЕКСІЙ", "OLEKSII"},
	{"олексій", "oleksii"},
	// Кк
	{"Київ", "Kyiv"},
	{"КИЇВ", "KYIV"},
	{"київ", "kyiv"},
	{"Коваленко", "Kovalenko"},
	{"КОВАЛЕНКО", "KOVALENKO"},
	{"коваленко", "kovalenko"},
	// Лл
	{"Лебедин", "Lebedyn"},
	{"ЛЕБЕДИН", "LEBEDYN"},
	{"лебедин", "lebedyn"},
	{"Леонід", "Leonid"},
	{"ЛЕОНІД", "LEONID"},
	{"леонід", "leonid"},
	// Мм
	{"Миколаїв", "Mykolaiv"},
	{"МИКОЛАЇВ", "MYKOLAIV"},
	{"миколаїв", "mykolaiv"},
	{"Маринич", "Marynych"},
	{"МАРИНИЧ", "MARYNYCH"},
	{"маринич", "marynych"},
	// Нн
	{"Ніжин", "Nizhyn"},
	{"НІЖИН", "NIZHYN"},
	{"ніжин", "nizhyn"},
	{"Наталія", "Nataliia"},
	{"НАТАЛІЯ", "NATALIIA"},
	{"наталія", "nataliia"},
	// Оо
	{"Одеса", "Odesa"},
	{"ОДЕСА", "ODESA"},
	{"одеса", "odesa"},
	{"Онищенко", "Onyshchenko"},
	{"ОНИЩЕНКО", "ONYSHCHENKO"},
	{"онищенко", "onyshchenko"},
	// Пп
	{"Полтава", "Poltava"},
	{"ПОЛТАВА", "POLTAVA"},
	{"полтава", "poltava"},
	{"Петро", "Petro"},
	{"ПЕТРО", "PETRO"},
	{"петро", "petro"},
	// Рр
	{"Решетилівка", "Reshetylivka"},
	{"РЕШЕТИЛІВКА", "RESHETYLIVKA"},
	{"решетилівка", "reshetylivka"},
	{"Рибчинський", "Rybchynskyi"},
	{"РИБЧИНСЬКИЙ", "RYBCHYNSKYI"},
	{"рибчинський", "rybchynskyi"},
	// Сс
	{"Суми", "Sumy"},
	{"СУМИ", "SUMY"},
	{"суми", "sumy"},
	{"Соломія", "Solomiia"},
	{"СОЛОМІЯ", "SOLOMIIA"},
	{"соломія", "solomiia"},
	// Тт
	{"Тернопіль", "Ternopil"},
	{"ТЕРНОПІЛЬ", "TERNOPIL"},
	{"тернопіль", "ternopil"},
	{"Троць", "Trots"},
	{"ТРОЦЬ", "TROTS"},
	{"троць", "trots"},
	// Уу
	{"Ужгород", "Uzhhorod"},
	{"УЖГОРОД", "UZHHOROD"},
	{"ужгород", "uzhhorod"},
	{"Уляна", "Uliana"},
	{"УЛЯНА", "ULIANA"},
	{"уляна", "uliana"},
	// Фф
	{"Фастів", "Fastiv"},
	{"ФАСТІВ", "FASTIV"},
	{"фастів", "fastiv"},
	{"Філіпчук", "Filipchuk"},
	{"ФІЛІПЧУК", "FILIPCHUK"},
	{"філіпчук", "filipchuk"},
	// Хх
	{"Харків", "Kharkiv"},
	{"ХАРКІВ", "KHARKIV"},
	{"харків", "kharkiv"},
	{"Христина", "Khrystyna"},
	{"ХРИСТИНА", "KHRYSTYNA"},
	{"христина", "khrystyna"},
	// Цц
	{"Біла Церква", "Bila Tserkva"},
	{"БІЛА ЦЕРКВА", "BILA TSERKVA"},
	{"біла церква", "bila tserkva"},
	{"Стеценко", "Stetsenko"},
	{"СТЕЦЕНКО", "STETSENKO"},
	{"стеценко", "stetsenko"},
	// Чч
	{"Чернівці", "Chernivtsi"},
	{"ЧЕРНІВЦІ", "CHERNIVTSI"},
	{"чернівці", "chernivtsi"},
	{"Шевченко", "Shevchenko"},
	{"ШЕВЧЕНКО", "SHEVCHENKO"},
	{"шевченко", "shevchenko"},
	// Шш
	{"Шостка", "Shostka"},
	{"ШОСТКА", "SHOSTKA"},
	{"шостка", "shostka"},
	{"Кишеньки", "Kyshenky"},
	{"КИШЕНЬКИ", "KYSHENKY"},
	{"кишеньки", "kyshenky"},
	// Щщ
	{"Щербухи", "Shcherbukhy"},
	{"ЩЕРБУХИ", "SHCHERBUKHY"},
	{"щербухи", "shcherbukhy"},
	{"Гоща", "Hoshcha"},
	{"ГОЩА", "HOSHCHA"},
	{"гоща", "hoshcha"},
	{"Гаращенко", "Harashchenko"},
	{"ГАРАЩЕНКО", "HARASHCHENKO"},
	{"гаращенко", "harashchenko"},
	// Юю
	{"Юрій", "Yurii"},
	{"ЮРІЙ", "YURII"},
	{"юрій", "yurii"},
	{"Корюківка", "Koriukivka"},
	{"КОРЮКІВКА", "KORIUKIVKA"},
	{"корюківка", "koriukivka"},
	// Яя
	{"Яготин", "Yahotyn"},
	{"ЯГОТИН", "YAHOTYN"},
	{"яготин", "yahotyn"},
	{"Ярошенко", "Yaroshenko"},
	{"ЯРОШЕНКО", "YAROSHENKO"},
	{"ярошенко", "yaroshenko"},
	{"Костянтин", "Kostiantyn"},
	{"КОСТЯНТИН", "KOSTIANTYN"},
	{"костянтин", "kostiantyn"},
	{"Знамʼянка", "Znamianka"},
	{"ЗНАМʼЯНКА", "ZNAMIANKA"},
	{"знамʼянка", "znamianka"},
	{"Феодосія", "Feodosiia"},
	{"ФЕОДОСІЯ", "FEODOSIIA"},
	{"феодосія", "feodosiia"},

	// Other examples
	{"", ""},
	{"єнот", "yenot"},
	{"їжак", "yizhak"},
	{"йорж", "yorzh"},
	{"юшка", "yushka"},
	{"яблуко", "yabluko"},
	{"зГраЯ", "zGHraIA"},
	{"Добрий день", "Dobryi den"},
	{"Привіт світ!", "Pryvit svit!"},
	{"Вредний єнот", "Vrednyi yenot"},
	{"Сміливий їжак", "Smilyvyi yizhak"},
	{"Риба йорж", "Ryba yorzh"},
	{"Грибна юшка", "Hrybna yushka"},
	{"Смачне яблуко", "Smachne yabluko"},
	{"Слава Україні! Glory to Ukraine!", "Slava Ukraini! Glory to Ukraine!"},
	{"test тест test", "test test test"},
	{"1234567890", "1234567890"},
	{"test TEST 123", "test TEST 123"},
}
