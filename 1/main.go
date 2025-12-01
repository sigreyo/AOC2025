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

		if isLeftspin(scanner.Text()) {
			current = current - number
		} else {
			current = current + number
		}

		if current == 0 || strings.HasSuffix(strconv.Itoa(current), "00") {
			zeroHits++
			current = 0
		}
	}
	fmt.Println(zeroHits)
}

func isLeftspin(line string) bool {
	return string(line[0]) == "L"
}
