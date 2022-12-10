package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverlap(t *testing.T) {
	cases := []struct {
		name string
		a    []int
		b    []int
		res  bool
	}{
		{
			name: "no overlap",
			a:    sectionsToInt("2-4"),
			b:    sectionsToInt("6-8"),
		},
		{
			name: "overlap 2-8",
			a:    sectionsToInt("2-8"),
			b:    sectionsToInt("3-7"),
			res:  true,
		},
		{
			name: "overlap 6-6",
			a:    sectionsToInt("6-6"),
			b:    sectionsToInt("4-6"),
			res:  true,
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.res, overlap(test.a, test.b), test.name)
	}
}

func Test3b(t *testing.T) {
	cases := []struct {
		name     string
		in       []string
		priority int
	}{
		{
			name: "first",
			in: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			priority: 18,
		},
		{
			name: "second",
			in: []string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			priority: 52,
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.priority, findBadgePriority(test.in[0], test.in[1], test.in[2]), test.name)
	}
}

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
