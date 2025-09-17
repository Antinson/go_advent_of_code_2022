package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const file = "test.txt"

func createStacks(st int) *[][]string {
	newStacks := [][]string{}

	for i := 1; i <= st+1; i++ {
		newStacks = append(newStacks, []string{})
	}

	return &newStacks
}

func p1() {

	data, _ := os.Open(file)
	defer data.Close()

	scanner := bufio.NewScanner(data)

	arrangement := []string{}
	instructions := []string{}
	flag := 0

	for scanner.Scan() {
		value := scanner.Text()

		if value == "" {
			flag = 1
			continue
		}

		if flag == 0 {
			arrangement = append(arrangement, value+"\n")
		} else {
			instructions = append(instructions, value)
		}
	}

	stacks := strings.Split(arrangement[len(arrangement)-1], "   ")

	fmt.Println(createStacks(len(stacks)))

	fmt.Println(instructions)
}
