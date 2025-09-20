package main

import "fmt"

func p2() {
	stacks, instrs := parseInput()
	for _, ins := range instrs {
		makeMoveP2(ins, stacks)
	}
	fmt.Println("Part 2 top crates:", topMessage(stacks))
}

// Part 2: move as a block (preserves order)
func makeMoveP2(ins *instruction, stacks [][]string) {
	from := ins.currentStack
	to := ins.newStack
	n := len(stacks[from])
	if ins.amount > n {
		return
	}
	block := stacks[from][n-ins.amount:]       // slice top block
	stacks[from] = stacks[from][:n-ins.amount] // shrink source
	stacks[to] = append(stacks[to], block...)  // append preserves order
}
