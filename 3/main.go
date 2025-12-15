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

	tot := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tot += collectJolts(scanner.Text())
	}
	fmt.Println(tot)
}

func collectJolts(line string) int {
	arr := intSlice(line)

	max := 0
	for baseIndex, base := range arr {
		for cmpIndex, cmp := range arr {
			high := getHighestJolt(base, cmp)
			if high > max && cmpIndex > baseIndex {
				max = high
			}
		}
	}

	return max
}

func getHighestJolt(num1, num2 int) int {
	tot1 := str(num1) + str(num2)

	return integer(tot1)
}

func intSlice(line string) []int {
	lineArr := strings.SplitSeq(line, "")
	intArr := []int{}
	for v := range lineArr {
		intArr = append(intArr, integer(v))
	}

	return intArr
}

func str(num int) string {
	return fmt.Sprintf("%d", num)
}
func integer(num string) int {
	number, _ := strconv.Atoi(num)
	return number
}
