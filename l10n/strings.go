package l10n

func Strings(lang string) map[string]string {
	switch lang {
	case "hu":
		return phrasesHU
	default:
		return phrases
	}
}
