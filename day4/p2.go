package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func overlayTwo(data string) int {

	newData := strings.Split(data, ",")

	test1 := strings.Split(newData[0], "-")
	test2 := strings.Split(newData[1], "-")

	range1 := newRange(string(test1[0]), string(test1[1]))
	range2 := newRange(string(test2[0]), string(test2[1]))

	var minRange int

	if range1.start <= range2.start {
		minRange = 1
	} else {
		minRange = 2
	}

	switch minRange {
	case 1:
		for i := range1.start; i <= range1.end; i++ {
			if i == range2.start || i == range2.end {
				return 1
			}
		}
	case 2:
		for i := range2.start; i <= range2.end; i++ {
			if i == range1.start || i == range1.end {
				return 1
			}
		}
	}
	return 0
}

func p2() {
	data, _ := os.Open("input.txt")
	defer data.Close()

	scanner := bufio.NewScanner(data)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		result := overlayTwo(line)
		total += result
	}
	fmt.Println(total)
}
