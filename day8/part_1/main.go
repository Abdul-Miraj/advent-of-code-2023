package main

import (
	"fmt"
	"strings"

	"aoc2023/utils"
)

var input = "../input.txt"

func getInstructions(text string) []int {
	var instructions []int
	for _, c := range text {

		if string(c) == "R" {
			instructions = append(instructions, 1)
			continue
		}
		instructions = append(instructions, 0)

	}
	return instructions
}

func main() {
	scanner, file := utils.GetScanner(input)
	defer file.Close()

	scanner.Scan()

	instructions := getInstructions(scanner.Text())

	nodes := map[string][]string{}
	for scanner.Scan() {

		t := scanner.Text()

		if t == "" {
			continue
		}

		direction := strings.Split(t, "=")

		key := strings.Trim(direction[0], " ")

		for _, n := range strings.Split(strings.TrimSpace(direction[1])[1:len(direction[1])-2], ",") {
			nodes[key] = append(nodes[key], strings.TrimSpace(n))
		}

	}

	// fmt.Println(nodes)

	var index, count int
	currLoc := "AAA"
	for {

		if currLoc == "ZZZ" {
			break
		}

		if index == len(instructions) {
			index = 0
		}

		// fmt.Println(currLoc)
		// fmt.Println(nodes[currLoc])
		// fmt.Println(instructions[index])
		// fmt.Println("====")
		currLoc = nodes[currLoc][instructions[index]]
		index++
		count++

	}

	fmt.Println(count)
}
