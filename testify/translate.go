package main

const (
	pigLatinSuffix string = "ay"
)

func Translate(in string) string {
	return in[1:] + in[0:1] + pigLatinSuffix
}
