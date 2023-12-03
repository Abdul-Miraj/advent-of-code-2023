package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	gameFile             = "../game.txt"
	maxColorForValidGame = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

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

func main() {
	gameRe := regexp.MustCompile(`(?m)(\d+ (green|red|blue),?);?`)
	numAndColorRe := regexp.MustCompile(`(?m)(?P<num>\d+) (?P<color>(green|red|blue))`)
	gameNumRe := regexp.MustCompile(`(?m)^Game (?P<gameNum>\d+)`)

	file, err := os.Open(gameFile)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int

	for scanner.Scan() {

		line := scanner.Text()

		validGame := true
		for _, match := range gameRe.FindAllString(line, -1) {

			result := extractCaptureGroup(numAndColorRe, match)
			num, _ := strconv.Atoi(result["num"])
			if maxColorForValidGame[result["color"]] < num {
				validGame = false
				break
			}
		}

		if validGame {
			gameNum, _ := strconv.Atoi(extractCaptureGroup(gameNumRe, line)["gameNum"])
			total += gameNum
			fmt.Println(fmt.Sprintf("Valid game %d, total %d", gameNum, total))
		}
	}

	fmt.Println(total)
}
