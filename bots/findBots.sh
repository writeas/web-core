#!/bin/bash

#
# Generates a Go map containing all bots that have accessed Write.as
#

cat /var/log/$1 | grep -i bot | awk -F\" '{print $4}' | sort | uniq > bots.txt

rm bots.go

cat > bots.go << EOM
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

func IsBot(ua string) bool {
    if _, ok := bots[ua]; ok {
        return true
    }
    return false
}
EOM
