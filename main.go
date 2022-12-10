package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("1.input")
	if err != nil {
		panic(err)
	}

	elves := map[int]int{}

	count := 1
	maxCal := 0
	maxElv := 0
	var keys []int

	for _, calories := range strings.Split(string(data), "\n") {
		if calories == "" {
			keys = append(keys, count)
			count++
			continue
		}

		iCalories, err := strconv.Atoi(calories)
		if err != nil {
			panic(err)
		}

		elves[count] += iCalories

		if elves[count] > elves[maxElv] {
			maxCal = elves[count]
			maxElv = count
		}

	}

	fmt.Println(fmt.Sprintf("keys: %+v %T", keys, keys))

	sort.SliceStable(keys, func(i, j int) bool {
		return elves[keys[i]] > elves[keys[j]]
	})

	fmt.Println(fmt.Sprintf("elves: %+v %T", elves, elves))
	fmt.Println(fmt.Sprintf("keys[0:3]: %+v %T", keys[0:3], keys[0:3]))
	part2 := elves[keys[0]] + elves[keys[1]] + elves[keys[2]]
	fmt.Println(fmt.Sprintf("part2: %+v %T", part2, part2))

	fmt.Println(fmt.Sprintf("maxCal: %+v %T", maxCal, maxCal))
	fmt.Println(fmt.Sprintf("maxElv: %+v %T", maxElv, maxElv))
}
