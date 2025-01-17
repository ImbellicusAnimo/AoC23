package main

import (
	"github.com/imbellicusanimo/AdventOfCode_2023/day1"
	"github.com/imbellicusanimo/AdventOfCode_2023/day2"
	"github.com/imbellicusanimo/AdventOfCode_2023/day3"
	"log"
)

func main() {
	log.Println("-------> Day 1 <-------")
	sum := day1.Run(1)
	log.Println("Sum:", sum)

	sum = day1.Run(2)
	log.Println("Sum:", sum)

	log.Println("-------> Day 2 <-------")
	sum = day2.Run(1)
	log.Println("Sum:", sum)

	sum = day2.Run(2)
	log.Println("Sum:", sum)

	log.Println("-------> Day 3 <-------")
	sum = day3.Run(1)
	log.Println("Sum:", sum)

	sum = day3.Run(2)
	log.Println("Sum:", sum)
}
