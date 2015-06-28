package logger

import (
	"regexp"
)

var postReg = regexp.MustCompile("(.*(/|id=))[a-zA-Z0-9]{12,13}(.*)")
var tokenReg = regexp.MustCompile("(.*t=)[a-zA-Z0-9]{32}(.*)")

func ScrubID(uri string) string {
	curStr := postReg.ReplaceAllString(uri, "$1[scrubbed]$3")
	curStr = tokenReg.ReplaceAllString(curStr, "$1[scrubbed]$2")
	return curStr
}
