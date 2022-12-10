package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwo(t *testing.T) {
	cases := []struct {
		name string
		in   []string
		out  int
	}{
		{
			name: "win with paper",
			in:   []string{"A Y"},
			out:  8,
		},
		{
			name: "lost with rock",
			in:   []string{"B X"},
			out:  1,
		},
		{
			name: "scissors draw",
			in:   []string{"C Z"},
			out:  6,
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.out, two(test.in), test.name)
	}
}
