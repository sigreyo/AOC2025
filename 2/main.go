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
	rawIDs := []string{}
	for scanner.Scan() {
		rawIDs = append(rawIDs, strings.Split(scanner.Text(), ",")...)
	}

	ranges := buildRanges(rawIDs)
	invalidIDs := checkDuplicates(ranges)

	fmt.Println(invalidIDs)
}

func buildRanges(data []string) map[int][]int {
	ranges := map[int][]int{}
	for i, v := range data {
		if v != "" {
			start, end := splitToStartAndEnd(v)
			ranges[i] = getRange(start, end)
		}
	}

	return ranges
}

func splitToStartAndEnd(ids string) (int, int) {
	splitIDs := strings.Split(ids, "-")

	start, err := strconv.Atoi(splitIDs[0])
	if err != nil {
		fmt.Println(err)
	}

	end, err := strconv.Atoi(splitIDs[1])
	if err != nil {
		fmt.Println(err)
	}

	return start, end
}

func getRange(start, end int) []int {
	idRange := make([]int, 0, end-start)
	for i := start; i <= end; i++ {
		idRange = append(idRange, i)
	}

	return idRange
}

func checkDuplicates(data map[int][]int) int {
	invalidIDs := 0
	for _, slice := range data {
		for _, v := range slice {
			if checkDuplicateOccurences(v) {
				invalidIDs += v
			}
		}
	}

	return invalidIDs
}

func checkDuplicateOccurences(num int) bool {
	strNum := strconv.Itoa(num)

	halfword := len(strNum) / 2

	return strNum[:halfword] == strNum[halfword:]
}
