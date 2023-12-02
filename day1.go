package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

//Consider your entire calibration document. What is the sum of all of the calibration values?

func Day1() int {
	input := "input.txt"
	var nums []string
	var numbers int
	var answer int
	readFile, err := os.Open(input)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		var found []rune
		for _, i := range text {
			if unicode.IsDigit(i) {
				found = append(found, i)
			}

		}
		nums = append(nums, string(found))
	}

	for _, i := range nums {
		thisNumber := string(i[0]) + string(i[len(i)-1])

		numbers, err = strconv.Atoi(thisNumber)
		if err != nil {
			log.Panic(err)
		}
		answer = answer + numbers
	}

	readFile.Close()
	return answer
}

func Day1Part2() int {
	input := "input.txt"
	var nums []string
	var numbers int
	var answer int

	wordMap := map[string]string{

		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	readFile, err := os.Open(input)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()

		for k, ok := range wordMap {

			text = strings.ReplaceAll(text, k, ok)
		}
		var found []rune
		for _, i := range text {
			if unicode.IsDigit(i) {
				found = append(found, i)
			}

		}
		nums = append(nums, string(found))
	}

	for _, i := range nums {
		thisNumber := string(i[0]) + string(i[len(i)-1])

		numbers, err = strconv.Atoi(thisNumber)
		if err != nil {
			log.Panic(err)
		}
		answer = answer + numbers
	}

	readFile.Close()
	return answer
}

// keys := []string{

// 	"one",
// 	"two",
// 	"three",
// 	"four",
// 	"five",
// 	"six",
// 	"seven",
// 	"eight",
// 	"nine",
// }
