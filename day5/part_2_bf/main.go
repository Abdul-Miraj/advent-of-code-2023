package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2023/utils"
)

var input = "../input.txt"

type alamancTranslation struct {
	lower int
	upper int
	diff  int
}

func main() {
	scanner, file := utils.GetScanner(input)
	defer file.Close()

	scanner.Scan()
	unprocessedSeeds := strings.Split(scanner.Text(), " ")

	var seeds []int
	for i := 1; i < len(unprocessedSeeds); i += 2 {

		n, _ := utils.IsNumber(unprocessedSeeds[i])
		r, _ := utils.IsNumber(unprocessedSeeds[i+1])

		for y := 0; y < r; y++ {
			seeds = append(seeds, n+y)
		}
	}

	fmt.Println("Seeds setup")

	var alamanc [][]alamancTranslation
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "map") {
			alamanc = append(alamanc, []alamancTranslation{})
			continue
		}

		t := strings.Split(line, " ")

		source, _ := strconv.Atoi(t[0])
		dest, _ := strconv.Atoi(t[1])
		diff, _ := strconv.Atoi(t[2])

		alamanc[len(alamanc)-1] = append(alamanc[len(alamanc)-1], alamancTranslation{
			lower: dest,
			upper: dest + diff - 1,
			diff:  source - dest,
		})
	}

	fmt.Println("processing translations...")
	processedSeeds := utils.NewSet[int](seeds...)
	for _, translations := range alamanc {
		newTranslations := utils.NewSet[int]()
		for _, t := range translations {
			for _, s := range processedSeeds.Slice() {
				if t.lower <= s && s <= t.upper {
					newTranslations.Add(s + t.diff)
					processedSeeds.Delete(s)
				}
			}
		}
		// fmt.Println(newTranslations)
		// fmt.Println(processedSeeds)

		processedSeeds = processedSeeds.Union(newTranslations)
		// fmt.Println(processedSeeds)
		// fmt.Println("=========")
	}

	fmt.Println("finding min")
	min := -1
	for _, s := range processedSeeds.Slice() {
		if s < min || min == -1 {
			min = s
		}
	}
	fmt.Println(min)
}
