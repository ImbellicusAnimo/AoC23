package day3

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/**
 * Advent of Code 2023 - Day 3:
 * https://adventofcode.com/2023/day/3
 */

func Run(part int) int {
	// read input
	lines := readTextInput("day3/day3_input.txt")

	var sum int

	if part == 1 {
		sum = part1(&lines)
	} else if part == 2 {
		sum = part2()
	} else {
		log.Fatal("Invalid part number")
	}
	return sum
}

func part1(lines *[]string) int {
	log.Println("Running -> Day 3 / Part 1")
	var sum int
	unifySymbols(lines)
	
	return sum
}

func unifySymbols(lines *[]string) {
	// replace all symbols with #
	for i, line := range *lines {
		(*lines)[i] = strings.ReplaceAll(line, "@", "#")
		(*lines)[i] = strings.ReplaceAll(line, "-", "#")
		(*lines)[i] = strings.ReplaceAll(line, "%", "#")
		(*lines)[i] = strings.ReplaceAll(line, "$", "#")
		(*lines)[i] = strings.ReplaceAll(line, "&", "#")
		(*lines)[i] = strings.ReplaceAll(line, "!", "#")
		(*lines)[i] = strings.ReplaceAll(line, "*", "#")
		(*lines)[i] = strings.ReplaceAll(line, "+", "#")
		(*lines)[i] = strings.ReplaceAll(line, "=", "#")
		(*lines)[i] = strings.ReplaceAll(line, ":", "#")
		(*lines)[i] = strings.ReplaceAll(line, ";", "#")
		(*lines)[i] = strings.ReplaceAll(line, ",", "#")
		(*lines)[i] = strings.ReplaceAll(line, "/", "#")
	}
}

func part2() int {
	log.Println("Running -> Day 3 / Part 2")

	var sum int

	return sum
}

/**
 * readTextInput reads the input file and returns a slice of strings
 */
func readTextInput(fileName string) []string {
	// read file day1_input.txt
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// read file line by line
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
