package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPriority(t *testing.T) {
	cases := []struct {
		name     string
		priority int
	}{
		{
			name:     "vJrwpWtwJgWrhcsFMMfFFhFp",
			priority: 16,
		},
		{name: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", priority: 38},
		{name: "PmmdzqPrVvPwwTWBwg", priority: 42},
		{name: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", priority: 22},
		{name: "ttgJtRGJQctTZtZT", priority: 20},
		{name: "CrZsJsPPZsGzwwsLwLmpwMDw", priority: 19},
	}

	for _, test := range cases {
		assert.Equal(t, test.priority, findpriority(test.name), test.name)
	}
}

func TestTwoB(t *testing.T) {
	cases := []struct {
		name string
		in   []string
		out  int
	}{
		{
			name: "draw with rock",
			in:   []string{"A Y"},
			out:  4,
		},
		{
			name: "lost with papper",
			in:   []string{"B X"},
			out:  1,
		},
		{
			name: "win with scissors",
			in:   []string{"C Z"},
			out:  7,
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.out, twob(test.in), test.name)
	}
}

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
