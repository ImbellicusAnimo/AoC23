package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * Advent of Code 2023 - Day 1:
 * https://adventofcode.com/2023/day/1
 */
const (
	ONE   = "one"
	TWO   = "two"
	THREE = "three"
	FOUR  = "four"
	FIVE  = "five"
	SIX   = "six"
	SEVEN = "seven"
	EIGHT = "eight"
	NINE  = "nine"
)

func Run(part int) int {
	// read input
	lines := readTextInput("day1/day1_input.txt")
	var sum int

	if part == 1 {
		sum = part1(&lines)
	} else if part == 2 {
		sum = part2(&lines)
	} else {
		log.Fatal("Invalid part number")
	}
	return sum
}

func part1(lines *[]string) int {
	log.Println("Running -> Day 1 / Part 1")
	var sum int
	// get coordinates
	for _, line := range *lines {
		x, y := getNumCoordinates(line)
		//log.Println("Coordinates:", x, y)
		number := buildNumberFromCoords(x, y)
		//log.Println("Number:", number)
		sum += number
	}
	return sum
}

func part2(lines *[]string) int {
	log.Println("Running -> Day 1 / Part 2")
	var sum int
	// get coordinates
	for _, line := range *lines {
		x, y := getAlphaNumCoordinates(line)
		//log.Println("Coordinates:", x, y)
		number := buildNumberFromStringCoords(x, y)

		//log.Println("Number:", number)
		sum += number
	}
	return sum
}

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

func buildNumberFromCoords(x, y int) int {
	combString := string(rune(x)) + string(rune(y))
	numberFromString, _ := strconv.Atoi(combString)
	return numberFromString
}

func buildNumberFromStringCoords(x, y string) int {
	combString := x + y
	numberFromString, _ := strconv.Atoi(combString)
	return numberFromString
}

func getNumCoordinates(input string) (int, int) {
	firstNumber := strings.IndexAny(input, "0123456789")
	secondNumber := strings.LastIndexAny(input, "0123456789")
	return int(input[firstNumber]), int(input[secondNumber])
}

func getAlphaNumCoordinates(input string) (string, string) {
	var firstIndex, lastIndex int
	var firstNumber, secondNumber string
	firstIndexOfNumber := strings.IndexAny(input, "0123456789")
	secondIndexOfNumber := strings.LastIndexAny(input, "0123456789")
	firstIndexOfOne := strings.Index(input, ONE)
	firstIndexOfTwo := strings.Index(input, TWO)
	firstIndexOfThree := strings.Index(input, THREE)
	firstIndexOfFour := strings.Index(input, FOUR)
	firstIndexOfFive := strings.Index(input, FIVE)
	firstIndexOfSix := strings.Index(input, SIX)
	firstIndexOfSeven := strings.Index(input, SEVEN)
	firstIndexOfEight := strings.Index(input, EIGHT)
	firstIndexOfNine := strings.Index(input, NINE)
	lastIndexOfOne := strings.LastIndex(input, ONE)
	lastIndexOfTwo := strings.LastIndex(input, TWO)
	lastIndexOfThree := strings.LastIndex(input, THREE)
	lastIndexOfFour := strings.LastIndex(input, FOUR)
	lastIndexOfFive := strings.LastIndex(input, FIVE)
	lastIndexOfSix := strings.LastIndex(input, SIX)
	lastIndexOfSeven := strings.LastIndex(input, SEVEN)
	lastIndexOfEight := strings.LastIndex(input, EIGHT)
	lastIndexOfNine := strings.LastIndex(input, NINE)

	firstIndex = firstIndexOfNumber
	if firstIndex != -1 {
		firstNumber = string(input[firstIndex])
	}
	if firstIndexOfOne != -1 && firstIndexOfOne < firstIndex {
		firstIndex = firstIndexOfOne
		firstNumber = "1"
	}
	if firstIndexOfTwo != -1 && firstIndexOfTwo < firstIndex {
		firstIndex = firstIndexOfTwo
		firstNumber = "2"
	}
	if firstIndexOfThree != -1 && firstIndexOfThree < firstIndex {
		firstIndex = firstIndexOfThree
		firstNumber = "3"
	}
	if firstIndexOfFour != -1 && firstIndexOfFour < firstIndex {
		firstIndex = firstIndexOfFour
		firstNumber = "4"
	}
	if firstIndexOfFive != -1 && firstIndexOfFive < firstIndex {
		firstIndex = firstIndexOfFive
		firstNumber = "5"
	}
	if firstIndexOfSix != -1 && firstIndexOfSix < firstIndex {
		firstIndex = firstIndexOfSix
		firstNumber = "6"
	}
	if firstIndexOfSeven != -1 && firstIndexOfSeven < firstIndex {
		firstIndex = firstIndexOfSeven
		firstNumber = "7"
	}
	if firstIndexOfEight != -1 && firstIndexOfEight < firstIndex {
		firstIndex = firstIndexOfEight
		firstNumber = "8"
	}
	if firstIndexOfNine != -1 && firstIndexOfNine < firstIndex {
		firstIndex = firstIndexOfNine
		firstNumber = "9"
	}

	lastIndex = secondIndexOfNumber
	secondNumber = string(input[lastIndex])
	if lastIndexOfOne != -1 && lastIndexOfOne > lastIndex {
		lastIndex = lastIndexOfOne
		secondNumber = "1"
	}
	if lastIndexOfTwo != -1 && lastIndexOfTwo > lastIndex {
		lastIndex = lastIndexOfTwo
		secondNumber = "2"
	}
	if lastIndexOfThree != -1 && lastIndexOfThree > lastIndex {
		lastIndex = lastIndexOfThree
		secondNumber = "3"
	}
	if lastIndexOfFour != -1 && lastIndexOfFour > lastIndex {
		lastIndex = lastIndexOfFour
		secondNumber = "4"
	}
	if lastIndexOfFive != -1 && lastIndexOfFive > lastIndex {
		lastIndex = lastIndexOfFive
		secondNumber = "5"
	}
	if lastIndexOfSix != -1 && lastIndexOfSix > lastIndex {
		lastIndex = lastIndexOfSix
		secondNumber = "6"
	}
	if lastIndexOfSeven != -1 && lastIndexOfSeven > lastIndex {
		lastIndex = lastIndexOfSeven
		secondNumber = "7"
	}
	if lastIndexOfEight != -1 && lastIndexOfEight > lastIndex {
		lastIndex = lastIndexOfEight
		secondNumber = "8"
	}
	if lastIndexOfNine != -1 && lastIndexOfNine > lastIndex {
		lastIndex = lastIndexOfNine
		secondNumber = "9"
	}
	return firstNumber, secondNumber
}
