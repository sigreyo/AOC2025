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
			if hasDuplicateOccurences(v) {
				invalidIDs += v
			}
		}
	}

	return invalidIDs
}

func hasDuplicateOccurences(num int) bool {
	strNum := strconv.Itoa(num)

	halfword := len(strNum) / 2

	if strNum[:halfword] == strNum[halfword:] {
		return true
	}

	return splitAndCheckForRepeats(strNum)
}

func splitAndCheckForRepeats(fullWord string) bool {
	if len(fullWord) == 1 {
		return false
	}

	if strings.Count(fullWord, string(fullWord[0])) == len(fullWord) {
		return true
	}

	testWordLen := 0
	if len(fullWord)%3 == 0 {
		testWordLen = len(fullWord) / 3
	}

	for i := 0; i < testWordLen; i++ {
		testWord := fullWord[testWordLen*i : testWordLen*(i+1)]
		if strings.Count(fullWord, testWord) > 2 {
			return true
		}
	}
	testWordLen = len(fullWord) / 5
	if testWordLen == 1 {
		testWordLen = 2
	}
	for i := 0; i < testWordLen; i++ {
		testWord := fullWord[testWordLen*i : testWordLen*(i+1)]
		if strings.Count(fullWord, testWord) > 4 {
			return true
		}
	}

	return false
}
