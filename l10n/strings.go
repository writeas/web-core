package l10n

// Strings returns a translation set that will take any term and return its
// translation.
func Strings(lang string) map[string]string {
	switch lang {
	case "hu":
		return phrasesHU
	default:
		return phrases
	}
}
