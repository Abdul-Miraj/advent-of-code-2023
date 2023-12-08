package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aoc2023/utils"
)

var input = "../input.txt"

const (
	HIGHCARD     = 0
	ONEPAIR      = 1
	TWOPAIR      = 2
	THREEOFAKIND = 3
	FULLHOUSE    = 4
	FOUROFAKIND  = 5
	FIVEOFAKIND  = 6
)

var deck = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type hand struct {
	bet   int
	cType int
	cards []int
}

func calculateCardType(hand string) (int, []int) {
	occ := map[rune]int{}

	var h []int
	for _, c := range hand {
		occ[c] += 1
		h = append(h, deck[string(c)])
	}
	var cType int
	switch len(occ) {
	case 5:
		cType = HIGHCARD
	case 4:
		cType = ONEPAIR
	case 3:
		cType = TWOPAIR
		for _, v := range occ {
			if v == 3 {
				cType = THREEOFAKIND
				break
			}
		}
	case 2:
		cType = FULLHOUSE
		for _, v := range occ {
			if v == 4 || v == 1 {
				cType = FOUROFAKIND
			}
		}
	case 1:
		cType = FIVEOFAKIND
	}

	return cType, h
}

func calculateTotalWinnings(hands []hand) int {
	var total int
	for i, h := range hands {
		total += ((i + 1) * h.bet)
	}

	return total
}

func main() {
	scanner, file := utils.GetScanner(input)
	defer file.Close()

	var hands []hand
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		b, _ := strconv.Atoi(line[1])

		cType, cards := calculateCardType(line[0])

		h := hand{
			bet:   b,
			cType: cType,
			cards: cards,
		}
		hands = append(hands, h)

	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].cType == hands[j].cType {
			for y := 0; y < 5; y++ {
				if hands[i].cards[y] < hands[j].cards[y] {
					return true
				} else if hands[i].cards[y] > hands[j].cards[y] {
					return false
				}
			}
		}

		return hands[i].cType < hands[j].cType
	})

	fmt.Println(calculateTotalWinnings(hands))
}
