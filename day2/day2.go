package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/**
 * Advent of Code 2023 - Day 2:
 * https://adventofcode.com/2023/day/2
 */

type Game struct {
	Id   int
	Sets []GameSet
}

type GameSet struct {
	Blue  int
	Red   int
	Green int
}

var games []Game

func Run(part int) int {
	// read input
	lines := readTextInput("day2/day2_input.txt")
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

// Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes,
// and 14 blue cubes. What is the sum of the IDs of those games?
func part1(lines *[]string) int {
	log.Println("Running -> Day 2 / Part 1")
	loadGames(lines)
	var sum int
	for _, game := range games {
		if checkGame(game) {
			log.Println("Game", game.Id, "is possible")
			sum += game.Id
		}
	}
	return sum
}

func checkGame(game Game) bool {
	// check if game is possible

	for _, set := range game.Sets {

		// check red
		if set.Red > 0 {
			if set.Red > 12 {
				return false
			}
		}

		// check green
		if set.Green > 0 {
			if set.Green > 13 {
				return false
			}
		}

		// check blue
		if set.Blue > 0 {
			if set.Blue > 14 {
				return false
			}
		}

	}
	return true
}

/*
 * loadGames parses the input lines and loads the games into the games slice
 */
func loadGames(lines *[]string) int {
	var numGames int
	for _, line := range *lines {
		// Game 30: 4 green, 5 blue, 1 red; 19 red, 18 blue, 3 green; 18 red, 18 blue, 1 green; 5 green, 14 blue, 4 red; 4 red, 3 green, 18 blue; 6 blue, 3 green, 17 red
		pre, suf, _ := strings.Cut(line, ":")
		gameIdStr, _ := strings.CutPrefix(pre, "Game ") // 30
		gameId, _ := strconv.Atoi(gameIdStr)
		sets := strings.Split(suf, ";") // 4 green, 5 blue, 1 red
		var game Game
		game.Id = gameId
		for _, set := range sets {
			var gameSet GameSet
			colors := strings.Split(set, ",")
			for _, color := range colors {
				color = strings.TrimSpace(color)
				if strings.HasSuffix(color, "green") {
					color = strings.TrimSuffix(color, "green")
					color = strings.TrimSpace(color)
					gameSet.Green, _ = strconv.Atoi(color)
				} else if strings.HasSuffix(color, "blue") {
					color = strings.TrimSuffix(color, "blue")
					color = strings.TrimSpace(color)
					gameSet.Blue, _ = strconv.Atoi(color)
				} else if strings.HasSuffix(color, "red") {
					color = strings.TrimSuffix(color, "red")
					color = strings.TrimSpace(color)
					gameSet.Red, _ = strconv.Atoi(color)
				}
			}
			game.Sets = append(game.Sets, gameSet)
		}
		games = append(games, game)
		numGames++
	}
	return numGames
}

func part2(lines *[]string) int {
	return 0
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
