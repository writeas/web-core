package bots

import "testing"

func TestIsBot(t *testing.T) {
	tests := map[string]bool{
		"Twitterbot/1.0":                                                                            true,
		"http.rb/2.2.2 (Mastodon/1.6.0; +https://insolente.im/)":                                    true,
		"http.rb/2.2.2 (Mastodon/1.5.1; +https://mastodon.cloud/)":                                  true,
		"http.rb/2.2.2 (Mastodon/1.6.0rc5; +https://mastodon.sdf.org/)":                             true,
		"http.rb/3.2.0 (Mastodon/2.4.3; +https://qoto.org/)":                                        true,
		"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko":              false,
		"Mozilla/5.0 (compatible; Applebot/0.3; +http://www.apple.com/go/applebot)":                 true,
		"Mozilla/5.0 (compatible; archive.org_bot +http://www.archive.org/details/archive.org_bot)": true,
		"Mozilla/5.0 (compatible; AhrefsBot/5.2; +http://ahrefs.com/robot/)":                        true,
	}

	for ua, r := range tests {
		if IsBot(ua) != r {
			t.Errorf("Expected bot = %t on '%s'", r, ua)
		}
	}
}
