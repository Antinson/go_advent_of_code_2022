package main

import (
	"bufio"
	"fmt"
	"os"
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

	// stacks := strings.Split(arrangement[len(arrangement)-1], "   ")

	// newStacks := createStacks(len(stacks))

	//fmt.Println(arrangement)

	for thing := 0; thing < len(arrangement)-1; thing++ {

		rnge := len(arrangement[thing]) - 1

		for {

			if rnge <= 4 {
				val := arrangement[thing][0:rnge]
				fmt.Printf("Range: %v - %v \t value: %v\n", rnge-3, rnge, val)
				break
			}

			fmt.Println(arrangement[thing][rnge-3 : rnge])
			rnge -= 3
		}
		// fmt.Printf("%v", thing)
		// fmt.Printf("%v", arrangement[thing][0:3])
		// fmt.Printf("%v\n", arrangement[thing][4:7])
	}

	fmt.Println(instructions)
}
