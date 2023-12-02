package main

import (
	"github.com/imbellicusanimo/AdventOfCode_2023/day1"
	"log"
)

func main() {
	sum := day1.Run(1)
	log.Println("Sum:", sum)

	sum = day1.Run(2)
	log.Println("Sum:", sum)
}
