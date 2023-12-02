package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() int {
	log.Println("Runnung -> Day 1 / Part 1")
	// read input
	lines := readTextInput("day1/day1_input.txt")

	// get coordinates
	var sum int
	for _, line := range lines {
		x, y := getCoordinates(line)
		//log.Println("Coordinates:", x, y)
		number := buildNumberFromCoords(x, y)
		//log.Println("Number:", number)
		sum += number
	}
	return sum
}

// Path: day1/day1.go
// 1. read file day1_input.txt

func readTextInput(fileName string) []string {

	// read file day1_input.txt
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read file line by line
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func buildNumberFromCoords(x, y int) int {
	combString := string(x) + string(y)
	numberFromString, _ := strconv.Atoi(combString)
	return numberFromString
}

func getCoordinates(input string) (int, int) {
	firstNumber := strings.IndexAny(input, "0123456789")
	secondNumber := strings.LastIndexAny(input, "0123456789")
	return int(input[firstNumber]), int(input[secondNumber])
}
