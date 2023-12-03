package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var gameFile = "../game.txt"

func extractCaptureGroup(r *regexp.Regexp, text string) map[string]string {
	p := r.FindStringSubmatch(text)
	result := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = p[i]
		}
	}

	return result
}

func powerOfSet(m map[string]int) int {
	total := 1
	for _, v := range m {
		total *= v
	}
	return total
}

func main() {
	gameRe := regexp.MustCompile(`(?m)(\d+ (green|red|blue),?);?`)
	numAndColorRe := regexp.MustCompile(`(?m)(?P<num>\d+) (?P<color>(green|red|blue))`)

	file, err := os.Open(gameFile)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int

	for scanner.Scan() {

		line := scanner.Text()

		minNumOfCubesForValidGame := map[string]int{}

		for _, match := range gameRe.FindAllString(line, -1) {

			result := extractCaptureGroup(numAndColorRe, match)
			num, _ := strconv.Atoi(result["num"])

			if num > minNumOfCubesForValidGame[result["color"]] {
				minNumOfCubesForValidGame[result["color"]] = num
			}

		}

		fmt.Println(minNumOfCubesForValidGame)

		total += powerOfSet(minNumOfCubesForValidGame)
	}

	fmt.Println(total)
}
