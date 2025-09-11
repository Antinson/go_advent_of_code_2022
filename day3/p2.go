package main

import (
	"bufio"
	"fmt"
	"os"
)

func splitAndGroupBuffer(three [3]string) int {

	valuesAre := [53]int{0: 0}
	total := 0
	count := 0

	for _, a := range three {
		for x := range a {
			tempVal := mapLettersa(int(a[x]))
			if valuesAre[tempVal] == count {
				valuesAre[tempVal] += 1
			}
		}
		count += 1
	}

	for v := range valuesAre {
		if valuesAre[v] == 3 {
			total += v
		}
	}

	fmt.Println(valuesAre)

	return total
}

func p2() {
	data, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	totalScore := 0

	buffer := [3]string{"", "", ""}
	count := 0

	for scanner.Scan() {

		line := scanner.Text()
		buffer[count] = line
		count += 1

		if count == 3 {
			totalScore += splitAndGroupBuffer(buffer)
			count = 0
		}
	}

	fmt.Println(totalScore)
}
