package l10n

// Strings returns a translation set that will take any term and return its
// translation.
func Strings(lang string) map[string]string {
	switch lang {
	case "ar":
		return phrasesAR
	case "de":
		return phrasesDE
	case "el":
		return phrasesEL
	case "es":
		return phrasesES
	case "fr":
		return phrasesFR
	case "hu":
		return phrasesHU
	case "it":
		return phrasesIT
	case "ja":
		return phrasesJA
	case "mk":
		return phrasesMK
	case "pl":
		return phrasesPL
	case "pt":
		return phrasesPT
	case "ro":
		return phrasesRO
	case "ru":
		return phrasesRU
	case "sv":
		return phrasesSV
	case "zh":
		return phrasesZH
	default:
		return phrases
	}
}
