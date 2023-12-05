package main

import (
	"fmt"
	"math"
	"strings"

	"aoc2023/utils"
)

var (
	input = "../input.txt"
	total int
)

func main() {
	scanner, file := utils.GetScanner(input)
	defer file.Close()

	for scanner.Scan() {
		t := scanner.Text()
		line := strings.Split(t, " ")

		winningNumbers := utils.NewSet[int]()
		ticketNumbers := utils.NewSet[int]()

		var isTicketNumber bool
		for _, l := range line {

			if l == "|" {
				isTicketNumber = true
			}
			v, ok := utils.IsNumber(l)
			if !ok {
				continue
			}

			if isTicketNumber {
				ticketNumbers.Add(v)
				continue
			}

			winningNumbers.Add(v)
		}

		intersection := winningNumbers.Intersection(ticketNumbers)
		fmt.Printf("Winning Numbers: %v | Ticket Numbers: %v\n", winningNumbers, ticketNumbers)
		fmt.Printf("Intersection: %v\n", intersection)

		total += int(math.Exp2(float64(intersection.Len() - 1)))
		fmt.Printf("Total: %d\n", total)

	}

	fmt.Println(total)
}
