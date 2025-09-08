package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calcScore(playerChoice string, opponentChoice string) int {
	total := 0

	mapNum := map[string]int{
		"A": 0, // Rock
		"B": 1, // Paper
		"C": 2, // Scissors
	}

	test := map[string][3]int{
		"X": {4, 1, 7}, // Rock
		"Y": {8, 5, 2}, // Paper
		"Z": {3, 9, 6}, // Scissors
	}

	total = test[playerChoice][mapNum[opponentChoice]]

	return total
}

func p1() {

	data, _ := os.Open("input.txt")

	defer data.Close()

	scanner := bufio.NewScanner(data)

	total := 0

	for scanner.Scan() {
		val := scanner.Text()

		tes := strings.Split(val, " ")
		tmpTotal := calcScore(tes[1], tes[0])
		fmt.Printf("Opp: %v \t Player: %v \t Score %v\n", tes[0], tes[1], tmpTotal)
		total += tmpTotal
	}

	fmt.Println(total)

}
