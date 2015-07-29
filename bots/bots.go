// This package helps the backend determine which clients are bots or crawlers.
// In Write.as, this is used to prevent certain things when viewing posts, like
// incrementing the view count.
package bots

var bots = map[string]bool{
	"bitlybot":                                                                            true,
	"crawlernutchtest/Nutch-1.9":                                                          true,
	"Domain Re-Animator Bot (http://domainreanimator.com) - support@domainreanimator.com": true,
	"ExactSeekCrawler/1.0":                                                                true,
	"Googlebot-Image/1.0":                                                                 true,
	"LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/3.1 +http://www.linkedin.com)":                                                                                                   true,
	"LivelapBot/0.2 (http://site.livelap.com/crawler)":                                                                                                                                                     true,
	"Mozilla/5.0 (compatible; AhrefsBot/5.0; +http://ahrefs.com/robot/)":                                                                                                                                   true,
	"Mozilla/5.0 (compatible; DotBot/1.1; http://www.opensiteexplorer.org/dotbot, help@moz.com)":                                                                                                           true,
	"Mozilla/5.0 (compatible; Exabot/3.0; +http://www.exabot.com/go/robot)":                                                                                                                                true,
	"Mozilla/5.0 (compatible; Findxbot/1.0; +http://www.findxbot.com)":                                                                                                                                     true,
	"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)":                                                                                                                             true,
	"Mozilla/5.0 (compatible; Kraken/0.1; http://linkfluence.net/; bot@linkfluence.net)":                                                                                                                   true,
	"Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/2.0; +http://go.mail.ru/help/robots)":                                                                                                              true,
	"Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/Robots/2.0; +http://go.mail.ru/help/robots)":                                                                                                       true,
	"Mozilla/5.0 (compatible; MojeekBot/0.6; +https://www.mojeek.com/bot.html)":                                                                                                                            true,
	"Mozilla/5.0 (compatible; OpenHoseBot/2.1; +http://www.openhose.org/bot.html)":                                                                                                                         true,
	"Mozilla/5.0 (compatible; OpsBot/1.0)":                                                                                                                                                                 true,
	"Mozilla/5.0 (compatible; PaperLiBot/2.1; http://support.paper.li/entries/20023257-what-is-paper-li)":                                                                                                  true,
	"Mozilla/5.0 (compatible; redditbot/1.0; +http://www.reddit.com/feedback)":                                                                                                                             true,
	"Mozilla/5.0 (compatible; SeznamBot/3.2; +http://fulltext.sblog.cz/)":                                                                                                                                  true,
	"Mozilla/5.0 (compatible; uMBot-LN/1.0; mailto: crawling@ubermetrics-technologies.com)":                                                                                                                true,
	"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)":                                                                                                                               true,
	"Mozilla/5.0 (compatible; woriobot +http://worio.com)":                                                                                                                                                 true,
	"Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)":                                                                                                                                     true,
	"Mozilla/5.0 (compatible; zitebot support [at] zite [dot] com +http://zite.com)":                                                                                                                       true,
	"Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)": true,
	"spiderbot":                                 true,
	"TelegramBot":                               true,
	"TelegramBot (like TwitterBot)":             true,
	"Traackr.com Bot":                           true,
	"Twitterbot/1.0":                            true,
	"voltron":                                   true,
	"Wotbox/2.01 (+http://www.wotbox.com/bot/)": true,
}

// IsBot returns whether or not the provided User-Agent string is a known bot
// or crawler.
func IsBot(ua string) bool {
	if _, ok := bots[ua]; ok {
		return true
	}
	return false
}