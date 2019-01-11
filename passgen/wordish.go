package passgen

import (
	"crypto/rand"
	"math/big"
)

var (
	ar   = []rune("aA4")
	cr   = []rune("cC")
	er   = []rune("eE3")
	fr   = []rune("fF")
	gr   = []rune("gG")
	hr   = []rune("hH")
	ir   = []rune("iI1")
	lr   = []rune("lL")
	nr   = []rune("nN")
	or   = []rune("oO0")
	rr   = []rune("rR")
	sr   = []rune("sS5")
	tr   = []rune("tT7")
	remr = []rune("bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ0123456789")
)

// NewWordish generates a password made of word-like words.
func NewWordish() string {
	b := []rune{}
	b = append(b, randLetter(cr))
	b = append(b, randLetter(hr))
	b = append(b, randLetter(ar))
	b = append(b, randLetter(nr))
	b = append(b, randLetter(gr))
	b = append(b, randLetter(er))
	b = append(b, randLetter(tr))
	b = append(b, randLetter(hr))
	b = append(b, randLetter(ir))
	b = append(b, randLetter(sr))
	b = append(b, randLetter(ar))
	b = append(b, randLetter(fr))
	b = append(b, randLetter(tr))
	b = append(b, randLetter(er))
	b = append(b, randLetter(rr))
	b = append(b, randLetter(lr))
	b = append(b, randLetter(or))
	b = append(b, randLetter(gr))
	b = append(b, randLetter(gr))
	b = append(b, randLetter(ir))
	b = append(b, randLetter(nr))
	b = append(b, randLetter(gr))
	b = append(b, randLetter(ir))
	b = append(b, randLetter(nr))
	for i := 0; i <= 7; i++ {
		b = append(b, randLetter(remr))
	}
	return string(b)
}

func randLetter(l []rune) rune {
	li, err := rand.Int(rand.Reader, big.NewInt(int64(len(l))))
	if err != nil {
		return rune(-1)
	}
	return l[li.Int64()]
}
