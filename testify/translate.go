package main

func Translate(in string) string {
	first := in[0:1]
	rest := in[1:]
	return rest + first + "ay"
}
