package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var textNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type numAndIndex struct {
	num   int
	index int
}

func main() {
	calibrationFile := "./calibrations.txt"

	file, err := os.Open(calibrationFile)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int

	for scanner.Scan() {

		line := scanner.Text()

		var first, last *numAndIndex
		for i, c := range line {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				continue
			}

			if first == nil {
				first = &numAndIndex{num: num, index: i}
			}

			last = &numAndIndex{num: num, index: i}
		}

		for word, num := range textNumbers {

			f := strings.Index(line, word)
			l := strings.LastIndex(line, word)

			if f != -1 && first.index > f {
				first = &numAndIndex{
					num:   num,
					index: f,
				}
			}

			if l != -1 && last.index < l {
				last = &numAndIndex{
					num:   num,
					index: l,
				}
			}

		}

		lineNum := strconv.Itoa(first.num) + strconv.Itoa(last.num)

		lineNumNum, _ := strconv.Atoi(lineNum)

		total += lineNumNum

		fmt.Println(fmt.Sprintf("Line: %s, First: %d, Last: %d, LineNum: %s, total: %d", line, *first, *last, lineNum, total))
	}

	fmt.Println(total)
}
