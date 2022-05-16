package wordchains

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleWordChains(t *testing.T) {
	w := &Words{}
	if err := w.GetEnglishWords(); err != nil {
		fmt.Println("Error reading words from file : " + err.Error())
		return
	}

	cases := []struct {
		inputStr1 string
		inputStr2 string
		output    string
	}{
		{"cat", "dog", "cat\ndat\ndot\ndog"},
		{"lead", "gold", "lead\nload\ngoad\ngold"},
	}

	for _, test := range cases {
		output := w.SimpleWordChains(test.inputStr1, test.inputStr2)
		assert.Equal(t, test.output, output)
	}
}
