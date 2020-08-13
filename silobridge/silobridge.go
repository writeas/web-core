package silobridge

// fakeAPInstances contains a list of sites that we allow writers to mention
// with the @handle@instance.tld syntax, plus the corresponding prefix to
// insert between `https://instance.tld/` and `handle` (e.g.
// https://medium.com/@handle)
var fakeAPInstances = map[string]string{
	"deviantart.com": "",
	"facebook.com":   "",
	"flickr.com":     "photos/",
	"github.com":     "",
	"instagram.com":  "",
	"medium.com":     "@",
	"reddit.com":     "user/",
	"twitter.com":    "",
	"wattpad.com":    "user/",
	"youtube.com":    "user/",
}

// Profile returns the full profile URL for a fake ActivityPub instance, based
// on the given handle and domain. If the domain isn't recognized, an empty
// string is returned.
func Profile(handle, domain string) string {
	prefix, ok := fakeAPInstances[domain]
	if !ok {
		return ""
	}
	return "https://" + domain + "/" + prefix + handle
}
