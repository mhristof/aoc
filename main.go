package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	for _, calories := range strings.Split(string(data), "\n") {
		fmt.Println(fmt.Sprintf("calories: %+v %T", calories, calories))

		if calories == "" {
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

	// dont forget to import "encoding/json"
	elvesJSON, err := json.MarshalIndent(elves, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(elvesJSON))

	fmt.Println(fmt.Sprintf("maxCal: %+v %T", maxCal, maxCal))
	fmt.Println(fmt.Sprintf("maxElv: %+v %T", maxElv, maxElv))
}
