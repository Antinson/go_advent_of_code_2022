package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func p1() {

	data, _ := os.Open("1.txt")

	defer data.Close()

	scanner := bufio.NewScanner(data)

	elfs := make([]Elf, 1)

	c := 0
	currentElf := &elfs[c]

	for scanner.Scan() {
		stringValue := scanner.Text()

		if stringValue == "" {
			elfs = append(elfs, Elf{})
			c++
			currentElf = &elfs[c]
		} else {
			tes, _ := strconv.Atoi(stringValue)
			currentElf.Total += tes
		}
	}

	highest := 0

	for i := range elfs {
		if elfs[i].Total > highest {
			highest = elfs[i].Total
		}
	}

	fmt.Println(highest)
}
