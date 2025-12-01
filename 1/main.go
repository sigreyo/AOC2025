package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	current := 50
	zeroHits := 0
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			fmt.Println(err)
		}

		line := buildLine(current, number, isLeftspin(scanner.Text()))
		current = line[len(line)-1]

		zeroes := checkZeroes(line)
		zeroHits += zeroes
	}

	fmt.Println(zeroHits)
}

func checkZeroes(line []int) int {
	zeroes := 0
	for _, v := range line {
		if v == 0 || strings.HasSuffix(strconv.Itoa(v), "00") {
			zeroes++
		}
	}

	return zeroes
}

func buildLine(start, spins int, isLeftspin bool) []int {
	slice := []int{}
	toAdd := start

	for range spins {
		if isLeftspin {
			toAdd--
		} else {
			toAdd++
		}

		if toAdd == 0 || strings.HasSuffix(strconv.Itoa(toAdd), "00") {
			toAdd = 0
		}

		slice = append(slice, toAdd)
	}

	return slice
}

func isLeftspin(line string) bool {
	return string(line[0]) == "L"
}
