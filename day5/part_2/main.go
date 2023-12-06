package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2023/utils"
)

var input = "../input.txt"

type seed struct {
	lower int
	upper int
}

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

	var seeds []seed

	for i := 1; i < len(unprocessedSeeds); i += 2 {
		n, _ := utils.IsNumber(unprocessedSeeds[i])
		r, _ := utils.IsNumber(unprocessedSeeds[i+1])

		seeds = append(seeds, seed{
			lower: n,
			upper: n + r - 1,
		})

	}

	fmt.Println(seeds)

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

	processedSeeds := utils.NewSet[seed](seeds...)
	for _, translations := range alamanc {
		newTranslations := utils.NewSet[seed]()
		partitionSeeds := utils.NewSet[seed]()
		for _, t := range translations {

			for _, s := range processedSeeds.Slice() {
				if t.lower <= s.lower && s.upper <= t.upper {
					partitionSeeds.Add(seed{
						lower: s.lower,
						upper: s.upper,
					})
					processedSeeds.Delete(s)
				} else if (t.lower <= s.lower && s.upper > t.upper) && (t.upper >= s.lower) {
					// 20-30
					// 50-60
					partitionSeeds.Add(seed{
						lower: s.lower,
						upper: t.upper,
					})

					partitionSeeds.Add(seed{
						lower: t.upper + 1,
						upper: s.upper,
					})
					processedSeeds.Delete(s)
				} else if (t.lower > s.lower && s.upper <= t.upper) && (t.lower <= s.upper) {
					// 50 - 60
					// 35 - 55
					// 35 - 49, 50-55

					partitionSeeds.Add(seed{
						lower: s.lower,
						upper: t.lower - 1,
					})

					partitionSeeds.Add(seed{
						lower: t.lower,
						upper: s.upper,
					})
					processedSeeds.Delete(s)
				}
			}

			fmt.Println(partitionSeeds)
			fmt.Println(processedSeeds)
			partitionSeeds = partitionSeeds.Union(processedSeeds)
			fmt.Println(partitionSeeds)
			fmt.Println("========")

			for _, s := range partitionSeeds.Slice() {
				if t.lower <= s.lower && s.upper <= t.upper {
					newTranslations.Add(seed{
						lower: s.lower + t.diff,
						upper: s.upper + t.diff,
					})
					partitionSeeds.Delete(s)
				}
			}
		}

		processedSeeds = partitionSeeds.Union(newTranslations)

	}

	min := -1
	for _, s := range processedSeeds.Slice() {
		if s.lower < min || min == -1 {
			min = s.lower
		}
	}
	fmt.Println(processedSeeds)
	fmt.Println(min)
}
