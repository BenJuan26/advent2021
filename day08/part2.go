package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2021"
)

const allLetters = "abcdefg"

func Subtract(p1, p2 string) string {
	for _, char := range p2 {
		p1 = strings.ReplaceAll(p1, string(char), "")
	}

	return p1
}

func ExtendReplacerArgs(args []string, randomLetter, trueLetter string) []string {
	if len(randomLetter) > 1 {
		panic(fmt.Sprintf("%s length: want 1 got %d", trueLetter, len(randomLetter)))
	}
	return append(args, randomLetter, trueLetter)
}

func GetSignalReplacer(patterns []string) *strings.Replacer {
	replacerArgs := []string{}
	var one, four, seven, eight string
	var sixes []string
	for _, pattern := range patterns {
		switch len(pattern) {
		case 2:
			one = pattern
		case 3:
			seven = pattern
		case 4:
			four = pattern
		case 6:
			sixes = append(sixes, pattern)
		case 7:
			eight = pattern
		}
	}

	// A
	a := Subtract(seven, one)
	replacerArgs = ExtendReplacerArgs(replacerArgs, a, "a")

	// B
	var missingSixes string
	for _, pattern := range sixes {
		m := Subtract(allLetters, pattern)
		if len(m) != 1 {
			panic(fmt.Sprintf("missingSixes result length: want got %d", len(m)))
		}
		missingSixes += m
	}

	b := Subtract(four, one)
	b = Subtract(b, missingSixes)
	replacerArgs = ExtendReplacerArgs(replacerArgs, b, "b")

	// D
	d := Subtract(four, one)
	d = Subtract(d, b)
	replacerArgs = ExtendReplacerArgs(replacerArgs, d, "d")

	// F
	f := Subtract(four, missingSixes)
	f = Subtract(f, b)
	replacerArgs = ExtendReplacerArgs(replacerArgs, f, "f")

	// C
	c := Subtract(four, b)
	c = Subtract(c, d)
	c = Subtract(c, f)
	replacerArgs = ExtendReplacerArgs(replacerArgs, c, "c")

	// E
	e := missingSixes
	e = Subtract(e, c)
	e = Subtract(e, d)
	replacerArgs = ExtendReplacerArgs(replacerArgs, e, "e")

	// G
	// probably a better way to do this, but it works
	g := eight
	g = Subtract(g, a)
	g = Subtract(g, b)
	g = Subtract(g, c)
	g = Subtract(g, d)
	g = Subtract(g, e)
	g = Subtract(g, f)

	replacerArgs = ExtendReplacerArgs(replacerArgs, g, "g")

	return strings.NewReplacer(replacerArgs...)
}

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	total := 0
	for _, line := range lines {
		lineSplit := strings.Split(line, " | ")

		patternsString := lineSplit[0]
		patterns := strings.Split(patternsString, " ")
		replacer := GetSignalReplacer(patterns)

		output := lineSplit[1]
		digits := strings.Split(output, " ")
		numbers := ""
		for _, pattern := range digits {
			sorted := []rune(pattern)
			convertedRunes := []rune(replacer.Replace(string(sorted)))
			sort.Slice(convertedRunes, func(i, j int) bool { return convertedRunes[i] < convertedRunes[j] })
			converted := string(convertedRunes)
			switch converted {
			case "abcefg":
				numbers += "0"
			case "cf":
				numbers += "1"
			case "acdeg":
				numbers += "2"
			case "acdfg":
				numbers += "3"
			case "bcdf":
				numbers += "4"
			case "abdfg":
				numbers += "5"
			case "abdefg":
				numbers += "6"
			case "acf":
				numbers += "7"
			case "abcdefg":
				numbers += "8"
			case "abcdfg":
				numbers += "9"
			default:
				panic("Unknown converted pattern " + converted)
			}
		}

		num, err := strconv.Atoi(numbers)
		if err != nil {
			panic(err)
		}

		total += num
	}

	fmt.Println(total)
}
