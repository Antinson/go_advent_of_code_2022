package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkPos(total int, highestArr []int) {
	highestArr[0] = total

	for x := 0; x < len(highestArr)-1; x++ {
		if highestArr[x] > highestArr[x+1] {
			temp := highestArr[x+1]
			highestArr[x+1] = highestArr[x]
			highestArr[x] = temp
		}
	}
}

func P2() {

	data, _ := os.Open("1.txt")

	defer data.Close()

	scanner := bufio.NewScanner(data)

	highestArr := [4]int{0, 0, 0, 0}
	total := 0

	for scanner.Scan() {
		stringValue := scanner.Text()

		if stringValue == "" {
			checkPos(total, highestArr[:])
			total = 0
		} else {
			val, _ := strconv.Atoi(stringValue)
			total += val
		}
	}

	checkPos(total, highestArr[:])

	sum := 0
	for i := 1; i < len(highestArr); i++ {
		sum += highestArr[i]
	}
	fmt.Println(highestArr)
	fmt.Println(sum)
}
