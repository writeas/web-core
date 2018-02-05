package bots

import "testing"

func TestIsBot(t *testing.T) {
	tests := map[string]bool{
		"Twitterbot/1.0":                                                               true,
		"http.rb/2.2.2 (Mastodon/1.6.0; +https://insolente.im/)":                       true,
		"http.rb/2.2.2 (Mastodon/1.5.1; +https://mastodon.cloud/)":                     true,
		"http.rb/2.2.2 (Mastodon/1.6.0rc5; +https://mastodon.sdf.org/)":                true,
		"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko": false,
	}

	for ua, r := range tests {
		if IsBot(ua) != r {
			t.Errorf("Expected bot = %t on '%s'", r, ua)
		}
	}
}
