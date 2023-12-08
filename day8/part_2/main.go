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

	var locations []string
	nodes := map[string][]string{}
	for scanner.Scan() {

		t := scanner.Text()

		if t == "" {
			continue
		}

		direction := strings.Split(t, "=")

		key := strings.Trim(direction[0], " ")

		if string(key[2]) == "A" {
			locations = append(locations, key)
		}

		for _, n := range strings.Split(strings.TrimSpace(direction[1])[1:len(direction[1])-2], ",") {
			nodes[key] = append(nodes[key], strings.TrimSpace(n))
		}

	}

	// fmt.Println(nodes)

	var reachedZ []int
	var loopCycle []int
	fmt.Println(locations)

	for _, l := range locations {
		var index, count int
		for {

			if string(l[2]) == "Z" {
				fmt.Println(l)
				reachedZ = append(reachedZ, count)
				var lc int
				for {

					if index == len(instructions) {
						index = 0
					}
					l = nodes[l][instructions[index]]
					index++
					lc++

					if string(l[2]) == "Z" {
						loopCycle = append(loopCycle, lc)
						break
					}

				}
				break
			}

			if index == len(instructions) {
				index = 0
			}

			l = nodes[l][instructions[index]]
			index++
			count++

		}
	}

	fmt.Println(reachedZ)
	fmt.Println(loopCycle)
}
