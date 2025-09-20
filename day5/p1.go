package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const file = "input.txt"

type instruction struct {
	amount       int
	currentStack int // 1-based, same as input
	newStack     int // 1-based, same as input
}

func createInstruction(parts []string) *instruction {
	// "move 3 from 2 to 1" -> ["move","3","from","2","to","1"]
	amt, _ := strconv.Atoi(parts[1])
	from, _ := strconv.Atoi(parts[3])
	to, _ := strconv.Atoi(parts[5])
	return &instruction{amount: amt, currentStack: from, newStack: to}
}

func printStacks(s [][]string) {
	for i := 1; i < len(s); i++ {
		fmt.Printf("%v\n", s[i])
	}
	fmt.Println()
}

func createStacks(n int) [][]string {
	// index 0 left empty so we can address stacks 1..n directly
	stacks := make([][]string, n+1)
	for i := 0; i <= n; i++ {
		stacks[i] = []string{}
	}
	return stacks
}

func splitRange(line string, stacks [][]string) {
	// read 3-chunks "[X]" or "   " separated by a space:
	// positions [0:3], [4:7], [8:11], ...
	start, end, idx := 0, 3, 1
	for {
		if end > len(line) {
			break
		}
		chunk := line[start:end]
		if chunk != "   " {
			// keep the token as "[X]" (AoC letters live at chunk[1])
			stacks[idx] = append(stacks[idx], chunk)
		}
		start = end + 1
		end += 4
		idx++
	}
}

func parseInput() ([][]string, []*instruction) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var arrangement []string
	var movesRaw []string
	seenBlank := false

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			seenBlank = true
			continue
		}
		if !seenBlank {
			arrangement = append(arrangement, line)
		} else {
			movesRaw = append(movesRaw, line)
		}
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	// Last line of arrangement is the index line: " 1   2   3 ..."
	indexLine := arrangement[len(arrangement)-1]
	numStacks := len(strings.Fields(indexLine))

	stacks := createStacks(numStacks)

	// Build stacks bottom-up (very important)
	for i := len(arrangement) - 2; i >= 0; i-- {
		splitRange(arrangement[i], stacks)
	}

	// Parse moves
	var instrs []*instruction
	for _, line := range movesRaw {
		parts := strings.Split(line, " ")
		instrs = append(instrs, createInstruction(parts))
	}
	return stacks, instrs
}

// Part 1: move one-by-one (reverses order)
func makeMoveP1(ins *instruction, stacks [][]string) {
	from := ins.currentStack
	to := ins.newStack
	if ins.amount > len(stacks[from]) {
		return
	}
	for c := 0; c < ins.amount; c++ {
		n := len(stacks[from])
		crate := stacks[from][n-1]             // top
		stacks[from] = stacks[from][:n-1]      // pop
		stacks[to] = append(stacks[to], crate) // push
	}
}

func topMessage(stacks [][]string) string {
	var b strings.Builder
	for i := 1; i < len(stacks); i++ {
		if len(stacks[i]) == 0 {
			continue
		}
		top := stacks[i][len(stacks[i])-1]
		// top looks like "[A]"; the letter is at index 1 if you want just the rune
		b.WriteByte(top[1])
	}
	return b.String()
}

func p1() {
	stacks, instrs := parseInput()
	for _, ins := range instrs {
		makeMoveP1(ins, stacks)
	}
	fmt.Println("Part 1 top crates:", topMessage(stacks))

}
