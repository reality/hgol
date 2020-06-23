package main

import (
	"fmt"

	"reality.rehab/hgol/board"
)

func main() {
	testString := "The ancient pond. A frog leaps in. Water's sound."
	//testString := "fuck the police. this is the life. i do not know."

	world := board.New(testString)

	generations := len(world.BinaryString)

	fmt.Print(generations)

	fmt.Print("\nInitial world:\n")

	world.String()

	for i := 0; i < generations; i++ {
		world.Progress()
	}

	fmt.Print("Final world:\n")
	world.String()

	fmt.Print(testString)
	fmt.Print('\n')

	world.Draw("basho.png")
}
