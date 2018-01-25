#!/bin/bash

#
# Generates a Go map containing all bots that have accessed Write.as from the
# application logs stored in /var/log/
#
# usage: findBots.sh application.log
#

cat /var/log/$1 | grep -i 'bot\|spider\|crawl\|scraper\|indexer\|voltron' | awk -F\" '{print $4}' | sort | uniq > bots.txt

rm bots.go

cat > bots.go << EOM
// This package helps the backend determine which clients are bots or crawlers.
// In Write.as, this is used to prevent certain things when viewing posts, like
// incrementing the view count.
package bots

var bots = map[string]bool {
EOM

while read b; do
    if [ -n "$b" ]; then
        echo "  \"$b\": true," >> bots.go
    fi
done <bots.txt

cat >> bots.go << EOM
};

// IsBot returns whether or not the provided User-Agent string is a known bot
// or crawler.
func IsBot(ua string) bool {
    if _, ok := bots[ua]; ok {
        return true
    }
    return false
}
EOM
