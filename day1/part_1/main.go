package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

		// loop string find ints

		var first, last *int
		for _, c := range line {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				continue
			}

			if first == nil {
				first = &num
			}

			last = &num

		}

		lineNum := strconv.Itoa(*first) + strconv.Itoa(*last)

		lineNumNum, _ := strconv.Atoi(lineNum)

		total += lineNumNum

		fmt.Println(fmt.Sprintf("Line: %s, First: %d, Last: %d, LineNum: %s, total: %d", line, *first, *last, lineNum, total))
	}

	fmt.Println(total)
}
