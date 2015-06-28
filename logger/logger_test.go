package logger

import (
	"testing"
)

type scrubTest struct {
	Input    string
	Expected string
}

var uris = []scrubTest{
	scrubTest{Input: "/1234567890123", Expected: "/[scrubbed]"},
	scrubTest{Input: "/acnsd8ndsklao", Expected: "/[scrubbed]"},
	scrubTest{Input: "/ACNSD8NDSKLAO", Expected: "/[scrubbed]"},
	scrubTest{Input: "/acNsD8NdSKlaO", Expected: "/[scrubbed]"},
	scrubTest{Input: "/acnsd8ndsklao.txt", Expected: "/[scrubbed].txt"},
	scrubTest{Input: "/8sj2kkjsn192.json", Expected: "/[scrubbed].json"},
	scrubTest{Input: "/acnsd8Ndsklao", Expected: "/[scrubbed]"},
	scrubTest{Input: "/12345678901", Expected: "/12345678901"},
	scrubTest{Input: "GET /8s9dja0vjbklj", Expected: "GET /[scrubbed]"},
	scrubTest{Input: "POST /8s9dja0vjbklj?delete=true", Expected: "POST /[scrubbed]?delete=true"},
	scrubTest{Input: "GET /8s9dja0vjbkl", Expected: "GET /[scrubbed]"},
	scrubTest{Input: "GET /asdf90as.txt", Expected: "GET /asdf90as.txt"},
	scrubTest{Input: "GET /api/999999999999", Expected: "GET /api/[scrubbed]"},
	scrubTest{Input: "DELETE /api/?id=8s9dja0vjbkl&t=123456789012345678901234567890ab", Expected: "DELETE /api/?id=[scrubbed]&t=[scrubbed]"},
	scrubTest{Input: "DELETE /api/8s9dja0vjbkl?t=123456789012345678901234567890ab", Expected: "DELETE /api/[scrubbed]?t=[scrubbed]"},
}

func TestScrubID(t *testing.T) {
	var scrubRes string

	for i := range uris {
		scrubRes = ScrubID(uris[i].Input)
		if scrubRes != uris[i].Expected {
			t.Errorf("#%d got %v, expected %v", i, scrubRes, uris[i].Expected)
		}
	}
}
