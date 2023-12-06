package main

import (
	"aoc2023/utils"
)

var input = "../input.txt"

func main() {
	scanner, file := utils.GetScanner(input)
	defer file.Close()

	for scanner.Scan() {
	}
}
