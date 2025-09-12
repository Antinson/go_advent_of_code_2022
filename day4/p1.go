package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rangeT struct {
	start int
	end   int
}

func newRange(startStr string, endStr string) *rangeT {
	start, _ := strconv.Atoi(startStr)
	end, _ := strconv.Atoi(endStr)

	return &rangeT{start: int(start), end: end}
}

func overlay(data string) int {

	result := 0

	newData := strings.Split(data, ",")

	range1 := newRange(string(newData[0][0]), string(newData[0][2]))
	range2 := newRange(string(newData[1][0]), string(newData[1][2]))

	if (range1.start <= range2.start && range1.end >= range2.end) ||
		(range2.start <= range1.start && range2.end >= range1.end) {
		result = 1
	}

	return result
}

func p1() {
	data, _ := os.Open("input.txt")
	defer data.Close()

	scanner := bufio.NewScanner(data)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		result := overlay(line)
		total += result
	}
	fmt.Println(total)
}
