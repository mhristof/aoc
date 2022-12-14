package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	res := fiveA(read("5.input"), map[int][]string{})
	fmt.Println(fmt.Sprintf("res: %+v %T", res, res))
}

func fiveA(data []string, stacks map[int][]string) string {
	if len(stacks) == 0 {
		stacks = map[int][]string{
			1: strings.Fields("S L W"),
			2: strings.Fields("J T N Q"),
			3: strings.Fields("S C H F J"),
			4: strings.Fields("T R M W N G B"),
			5: strings.Fields("T R L S D H Q B"),
			6: strings.Fields("M J B V F H R L"),
			7: strings.Fields("D W R N J M"),
			8: strings.Fields("B Z T F H N D J"),
			9: strings.Fields("H L Q N B F T"),
		}
	}
	for _, move := range data[10:] {
		fmt.Println(fmt.Sprintf("move: %+v %T", move, move))
		fields := strings.Fields(move)
		moveI := []int{atoi(fields[1]), atoi(fields[3]), atoi(fields[5])}
		stacks = moveStacks(stacks, moveI)
	}

	// dont forget to import "encoding/json"
	stacksJSON, err := json.MarshalIndent(stacks, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(stacksJSON))

	var ret []string
	for i := 1; i <= len(stacks); i++ {
		last := stacks[i][len(stacks[i])-1]
		ret = append(ret, last)
	}

	// fmt.Println(fmt.Sprintf("ret: %+v %T", ret, ret))

	return strings.Join(ret, "")
}

func moveStacks(stacks map[int][]string, moves []int) map[int][]string {
	for i := 0; i < moves[0]; i++ {
		srcLen := len(stacks[moves[1]])
		moved := stacks[moves[1]][srcLen-1]
		stacks[moves[1]] = stacks[moves[1]][:srcLen-1]
		stacks[moves[2]] = append(stacks[moves[2]], moved)
	}
	return stacks
}

func atoi(in string) int {
	number, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return number
}

func fourAb(data []string) int {
	var counter int

	for _, team := range data {
		fields := strings.Split(team, ",")
		if overlap(sectionsToInt(fields[0]), sectionsToInt(fields[1]), false) {
			counter++
		}
	}

	return counter
}

func fourA(data []string) int {
	var counter int

	for _, team := range data {
		fields := strings.Split(team, ",")
		if overlap(sectionsToInt(fields[0]), sectionsToInt(fields[1]), true) {
			counter++
		}
	}

	return counter
}

func overlap(a, b []int, all bool) bool {
	overlap := 0
	// fmt.Println(fmt.Sprintf("a: %+v %T", a, a))
	// fmt.Println(fmt.Sprintf("b: %+v %T", b, b))
	for _, first := range a {
		for _, second := range b {
			if first == second {
				overlap += 1
			}
		}
	}

	// fmt.Println(fmt.Sprintf("overlap: %+v %T", overlap, overlap))

	if all {
		return overlap >= len(a) || overlap >= len(b)
	}
	return overlap > 0
}

func sectionsToInt(section string) []int {
	fields := strings.Split(section, "-")

	start, err := strconv.Atoi(fields[0])
	if err != nil {
		panic(err)
	}

	stop, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}

	var ret []int
	for i := start; i <= stop; i++ {
		ret = append(ret, i)
	}
	return ret
}

func threeB(data []string) int {
	backpacks := make([]string, 3)
	prio := 0
	for i, line := range data {
		backpacks[i%3] = line
		if i%3 == 2 {
			prio += findBadgePriority(backpacks[0], backpacks[1], backpacks[2])
			fmt.Println(fmt.Sprintf("backpacks: %+v %T", backpacks, backpacks))
		}
	}
	return prio
}

func findBadgePriority(a, b, c string) int {
	for _, first := range a {
		for _, second := range b {
			for _, third := range c {
				if first == second && second == third {
					return runePriority(first)
				}
			}
		}
	}
	fmt.Println(fmt.Sprintf("a: %+v %T", a, a))
	fmt.Println(fmt.Sprintf("b: %+v %T", b, b))
	fmt.Println(fmt.Sprintf("c: %+v %T", c, c))
	panic("panic!")
}

func runePriority(in rune) int {
	if in >= 97 {
		return int(in) - 96
	}
	return int(in) - 64 + 26
}

func threeA(data []string) int {
	var sum int

	for _, backpack := range data {
		sum += findpriority(backpack)
	}

	return sum
}

func findpriority(backpack string) int {
	l := len(backpack)
	first := backpack[0 : l/2]
	secon := backpack[l/2 : l]
	fmt.Println(fmt.Sprintf("first: %+v %T", first, first))
	fmt.Println(fmt.Sprintf("secon: %+v %T", secon, secon))

	var extra rune

	for _, letter := range first {
		found := false
		for _, sLetter := range secon {
			if sLetter == letter {
				found = true

				break
			}
		}

		if found {
			fmt.Println(fmt.Sprintf("letter: %s %T", string(letter), letter))
			extra = letter

			break
		}

	}

	return runePriority(extra)
}

func twob(data []string) int {
	scores := map[string]int{
		"rock":     1,
		"papper":   2,
		"scissors": 3,
		"win":      6,
		"lost":     0,
		"draw":     3,
	}

	words := map[string]string{
		"X": "lost",
		"Y": "draw",
		"Z": "win",
		"A": "rock",
		"B": "papper",
		"C": "scissors",
	}

	score := 0
	for _, game := range data {
		fields := strings.Fields(game)

		them := words[fields[0]]
		result := words[fields[1]]

		oldScore := score
		me := ""

		switch {
		case result == "draw":
			me = them
		case result == "win" && them == "papper":
			me = "scissors"
		case result == "win" && them == "rock":
			me = "papper"
		case result == "win" && them == "scissors":
			me = "rock"
		case result == "lost" && them == "papper":
			me = "rock"
		case result == "lost" && them == "rock":
			me = "scissors"
		case result == "lost" && them == "scissors":
			me = "papper"
		}

		score += scores[me] + scores[result]

		fmt.Println(fmt.Sprintf("game: %s %s %s %d %d", them, me, result, score-oldScore, score))
	}
	fmt.Println(fmt.Sprintf("score: %+v %T", score, score))
	return score
}

func two(data []string) int {
	scores := map[string]int{
		"rock":     1,
		"papper":   2,
		"scissors": 3,
	}

	translate := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	words := map[string]string{
		"X": "rock",
		"Y": "papper",
		"Z": "scissors",
		"A": "rock",
		"B": "papper",
		"C": "scissors",
	}

	score := 0
	for _, game := range data {
		fields := strings.Fields(game)

		// fmt.Println(fmt.Sprintf("game: %+v %T", game, game))
		// fmt.Println(fmt.Sprintf("fields: %+v %T", fields, fields))

		them := words[translate[fields[0]]]
		me := words[fields[1]]

		oldScore := score
		res := ""

		switch {
		case them == me:
			score += scores[me] + 3
			res = "draw"
		case me == "papper" && them != "scissors":
			score += scores[me] + 6
			res = "win"
		case me == "rock" && them != "papper":
			score += scores[me] + 6
			res = "win"
		case me == "scissors" && them != "rock":
			score += scores[me] + 6
			res = "win"
		default:
			score += scores[me]
			res = "lost"
		}

		// fmt.Println(fmt.Sprintf("them: %+v %T", them, them))
		fmt.Println(fmt.Sprintf("game: %s %s %s %d %d", them, me, res, score-oldScore, score))
	}
	fmt.Println(fmt.Sprintf("score: %+v %T", score, score))
	return score
}

func read(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func part1() {
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
