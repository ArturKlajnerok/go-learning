package main

import (
	"strings"
)

const (
	pigLatinSuffix             string = "ay"
	firstLetterExceptions      string = "aeiou"
	firstLetterExceptionSuffix string = "d" + pigLatinSuffix
)

func Translate(in string) string {
	first := in[0:1]
	if strings.Contains(firstLetterExceptions, first) {
		return in + firstLetterExceptionSuffix
	} else {
		return in[1:] + first + pigLatinSuffix
	}
}
