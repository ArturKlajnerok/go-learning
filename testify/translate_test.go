package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var translateTests = []struct {
	in  string
	out string
}{
	{
		in:  "dog",
		out: "ogday",
	},
	{
		in:  "cat",
		out: "atcay",
	},
	{
		in:  "hat",
		out: "athay",
	},
}

func TestTranslate(t *testing.T) {
	for i, test := range translateTests {
		actual := Translate(test.in)
		assert.Equal(t, test.out, actual, "Test %d", i)
	}
}
