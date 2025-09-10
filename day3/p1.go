package main

import (
	"bufio"
	"fmt"
	"os"
)

func splitAndGroup(a string) {

	firstSack := a[:len(a)/2]
	//secondSack := a[:len(a)/2]

	for letter := range firstSack {
		fmt.Println(letter)
	}

	//fmt.Printf("%v \t\t %v\n", firstSack, secondSack)
}

func p1() {

	data, _ := os.Open("test.txt")
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		splitAndGroup(line)
	}

}
