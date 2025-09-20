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
			arrangement = append(arrangement, value)
		} else {
			instructions = append(instructions, value)
		}
	}

	stacks := strings.Split(arrangement[len(arrangement)-1], "   ")

	newStacks := createStacks(len(stacks))

	fmt.Println(arrangement)

	for thing := 0; thing < len(arrangement)-1; thing++ {

		v := strings.TrimSpace(arrangement[thing])
		splitRange(v, newStacks)

	}
	fmt.Println(newStacks)
	fmt.Println(instructions)
}

func splitRange(s string, stack *[][]string) {
	t := strings.Split(s, " ")
	c := len(t)

	for index, value := range t {
		(*stack)[index] = append((*stack)[index], value)
	}

	fmt.Println(c, t)
}
