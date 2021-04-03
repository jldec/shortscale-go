package shortscale

import (
	"strings"
)

// Shortscale converts numbers into English words.
// Supports positive integers from 0 to 999_999_999_999_999_999.
// Larger values return "(big number)".
func Shortscale(n uint64) string {
	if n <= 20 {
		return numwords[n]
	}
	if n > 999_999_999_999_999_999 {
		return "(big number)"
	}
	var b strings.Builder
	b.Grow(238)
	writeScale(&b, n, 1_000_000_000_000_000) // quadrillions
	writeScale(&b, n, 1_000_000_000_000)     // trillions
	writeScale(&b, n, 1_000_000_000)         // billions
	writeScale(&b, n, 1_000_000)             // millions
	writeScale(&b, n, 1_000)                 // thousands
	writeHundreds(&b, n)
	writeTensAndUnits(&b, n, b.Len() > 0)
	return b.String()
}

func writeTensAndUnits(b *strings.Builder, n uint64, ifAnd bool) {
	n = n % 100
	if n == 0 {
		return
	}
	if ifAnd {
		writeWord(b, "and")
	}
	if n <= 20 {
		writeWord(b, numwords[n])
		return
	}
	writeWord(b, numwords[n/10*10]) // tens
	units := n % 10
	if units > 0 {
		writeWord(b, numwords[units])
	}
}

func writeHundreds(b *strings.Builder, n uint64) {
	n = n / 100 % 10
	if n == 0 {
		return
	}
	writeWord(b, numwords[n])
	writeWord(b, numwords[100])
}

func writeScale(b *strings.Builder, n uint64, thousands uint64) {
	n = n / thousands % 1_000
	if n == 0 {
		return
	}
	writeHundreds(b, n)
	writeTensAndUnits(b, n, (n/100%10) > 0)
	writeWord(b, numwords[thousands])
}

func writeWord(b *strings.Builder, word string) {
	if b.Len() > 0 {
		b.WriteString(" ")
	}
	b.WriteString(word)
}

var numwords = map[uint64]string{
	0:                     "zero",
	1:                     "one",
	2:                     "two",
	3:                     "three",
	4:                     "four",
	5:                     "five",
	6:                     "six",
	7:                     "seven",
	8:                     "eight",
	9:                     "nine",
	10:                    "ten",
	11:                    "eleven",
	12:                    "twelve",
	13:                    "thirteen",
	14:                    "fourteen",
	15:                    "fifteen",
	16:                    "sixteen",
	17:                    "seventeen",
	18:                    "eighteen",
	19:                    "nineteen",
	20:                    "twenty",
	30:                    "thirty",
	40:                    "forty",
	50:                    "fifty",
	60:                    "sixty",
	70:                    "seventy",
	80:                    "eighty",
	90:                    "ninety",
	100:                   "hundred",
	1_000:                 "thousand",
	1_000_000:             "million",
	1_000_000_000:         "billion",
	1_000_000_000_000:     "trillion",
	1_000_000_000_000_000: "quadrillion",
}
