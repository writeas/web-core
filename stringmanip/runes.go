package stringmanip

/*
	The MIT License (MIT)

	Copyright (c) 2016 Sergey Kamardin

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/
// Source: https://github.com/gobwas/glob/blob/master/util/runes/runes.go

func IndexRune(str string, r rune) int {
	s := []rune(str)
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}

func LastIndexRune(str string, needle rune) int {
	s := []rune(str)
	needles := []rune{needle}
	ls, ln := len(s), len(needles)

	switch {
	case ln == 0:
		if ls == 0 {
			return 0
		}
		return ls
	case ln == 1:
		return IndexLastRune(s, needles[0])
	case ln == ls:
		if EqualRunes(s, needles) {
			return 0
		}
		return -1
	case ln > ls:
		return -1
	}

head:
	for i := ls - 1; i >= 0 && i >= ln; i-- {
		for y := ln - 1; y >= 0; y-- {
			if s[i-(ln-y-1)] != needles[y] {
				continue head
			}
		}

		return i - ln + 1
	}

	return -1
}

func IndexLastRune(s []rune, r rune) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == r {
			return i
		}
	}

	return -1
}

func EqualRunes(a, b []rune) bool {
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}

		return true
	}

	return false
}
