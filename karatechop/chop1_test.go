package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMediumWordChains(t *testing.T) {

	cases := []struct {
		numberToFind  int
		listOfNumbers []int
		output        int
	}{
		{9, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8},
		{3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2},
	}

	for _, test := range cases {
		output := chop1(test.numberToFind, test.listOfNumbers)
		assert.Equal(t, test.output, output)
	}
}
