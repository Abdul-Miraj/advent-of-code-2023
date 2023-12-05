package main

import (
	"fmt"
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

	cardNumberToNumOfCards := map[int]int{}
	var cardNumber int
	for scanner.Scan() {
		cardNumber += 1
		cardNumberToNumOfCards[cardNumber] += 1
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

		winners := intersection.Len()

		for i := 1; i < winners+1; i++ {
			cardNumberToNumOfCards[cardNumber+i] += cardNumberToNumOfCards[cardNumber]
		}

	}

	for _, v := range cardNumberToNumOfCards {
		total += v
	}
	fmt.Println(total)
}
