package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calcScoreTwo(opponentChoice string, playerChoice string) int {
	total := 0

	mapNum := map[string][3]int{
		"A": {1, 4, 7}, // Rock
		"B": {2, 5, 8}, // Paper
		"C": {3, 6, 9}, // Scissors
	}

	switch playerChoice {
	case "X": // Need to loose
		total = mapNum[opponentChoice][0]
	case "Y": // Need to draw
		total = mapNum[opponentChoice][1]
	case "Z": // Need to win
		total = mapNum[opponentChoice][2]

	}

	return total
}

func p2() {

	data, _ := os.Open("input_test.txt")

	defer data.Close()

	scanner := bufio.NewScanner(data)

	total := 0

	for scanner.Scan() {
		val := scanner.Text()

		tes := strings.Split(val, " ")
		tmpTotal := calcScoreTwo(tes[0], tes[1])
		fmt.Printf("Opp: %v \t Player: %v \t Score %v\n", tes[0], tes[1], tmpTotal)
		total += tmpTotal
	}

	fmt.Println(total)

}
