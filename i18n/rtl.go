package i18n

var rtlLangs = map[string]bool{
	"ar": true, // Arabic
	"dv": true, // Divehi
	"fa": true, // Persian (Farsi)
	"ha": true, // Hausa
	"he": true, // Hebrew
	"iw": true, // Hebrew (old code)
	"ji": true, // Yiddish (old code)
	"ps": true, // Pashto, Pushto
	"ur": true, // Urdu
	"yi": true, // Yiddish
}

func LangIsRTL(lang string) bool {
	if _, ok := rtlLangs[lang]; ok {
		return true
	}
	return false
}
