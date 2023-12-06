package main

import (
	"fmt"
	"strings"

	"aoc2023/utils"
)

var input = "../input.txt"

type race struct {
	time     int
	distance int
}

func (r race) calculateDifferentWaysToBeat() int {
	var res int
	for i := 1; i < r.time; i++ {
		if i*(r.time-i) > r.distance {
			res += 1
		}
	}
	return res
}

func filterNonDigits(s []string) []int {
	var res []int
	for _, t := range s {
		v, ok := utils.IsNumber(t)
		if !ok {
			continue
		}

		res = append(res, v)
	}

	return res
}

func main() {
	scanner, file := utils.GetScanner(input)
	defer file.Close()

	scanner.Scan()
	uTimes, _ := strings.CutPrefix(scanner.Text(), "Time:")

	scanner.Scan()
	uDistances, _ := strings.CutPrefix(scanner.Text(), "Distance:")

	times := filterNonDigits(strings.Split(uTimes, " "))
	distances := filterNonDigits(strings.Split(uDistances, " "))

	var races []race
	for i := 0; i < len(times); i++ {
		races = append(races, race{
			time:     times[i],
			distance: distances[i],
		})
	}

	total := 1
	for _, r := range races {
		total *= r.calculateDifferentWaysToBeat()
	}

	fmt.Println(total)
}
