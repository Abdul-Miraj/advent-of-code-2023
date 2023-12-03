package main

import (
	"fmt"
	"strconv"

	"aoc2023/utils"
)

var input = "../input.txt"

type partLocation struct {
	row int
	col int
}

var (
	total   int
	visited = utils.NewSet[[2]int]()
)

func main() {
	scanner, file := utils.GetScanner(input)
	defer file.Close()

	var schematic [][]string
	var pLoc []partLocation
	for scanner.Scan() {
		line := scanner.Text()

		schematic = append(schematic, []string{})
		for _, c := range line {
			symbol := string(c)
			schematic[len(schematic)-1] = append(schematic[len(schematic)-1], symbol)

			if symbol != "." {
				if _, err := strconv.Atoi(symbol); err != nil {
					partLocation := partLocation{
						row: len(schematic) - 1,
						col: len(schematic[len(schematic)-1]) - 1,
					}
					pLoc = append(pLoc, partLocation)

				}
			}
		}
	}

	for _, loc := range pLoc {
		checkNeighbors(loc.row, loc.col, schematic)
		// fmt.Println("=====")
	}

	// fmt.Println(visited)
	fmt.Println(total)
}

func checkNeighbors(row int, col int, schematic [][]string) int {
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	for _, d := range directions {
		cRow := row + d[0]
		cCol := col + d[1]

		if visited.Exists([2]int{cRow, cCol}) {
			continue
		}

		if cRow < 0 || cRow >= len(schematic) || cCol < 0 || cCol >= len(schematic[cRow]) {
			continue
		}

		symbol := schematic[cRow][cCol]

		_, err := strconv.Atoi(symbol)
		if err != nil {
			continue
		}

		// fmt.Println(fmt.Sprintf("%d, %d, %s", cRow, cCol, symbol))
		n := scanForNumber(cRow, cCol, schematic)

		total += n

		fmt.Printf("Symbol: %s, Num: %d, Total %d\n", schematic[row][col], n, total)
	}

	return total
}

func scanForNumber(row, col int, schematic [][]string) int {
	for col > -1 {
		_, err := strconv.Atoi(schematic[row][col])
		if err != nil {
			col++
			break
		}

		col--
	}

	var potNum string

	if col == -1 {
		col = 0
	}
	// fmt.Println(row, col)
	for col < len(schematic[row]) {

		_, err := strconv.Atoi(schematic[row][col])
		if err != nil {
			break
		}

		potNum += schematic[row][col]
		visited.Add([2]int{row, col})
		col++
	}

	num, _ := strconv.Atoi(potNum)

	return num
}

// 553584
