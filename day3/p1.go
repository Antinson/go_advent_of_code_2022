package main

import (
	"bufio"
	"fmt"
	"os"
)

func mapLettersa(letter int) int {
	if letter >= 97 {
		return letter - 96
	} else {
		return letter - 64 + 26
	}
}

func splitAndGroup(a string) int {

	firstSack := a[0 : len(a)/2]
	secondSack := a[len(a)/2:]

	checkValues := make(map[int]int, 5)

	for i, value := range firstSack {
		for j := range secondSack {
			if firstSack[i] == secondSack[j] {
				checkValues[int(value)] = mapLettersa(int(value))
			}
		}
	}

	total := 0

	for _, value := range checkValues {
		total += int(value)
	}

	return total

}

func p1() {

	data, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		totalScore += splitAndGroup(line)
	}

	fmt.Println(totalScore)

}
