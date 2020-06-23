package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"reality.rehab/hgol/board"
)

func main() {
	//testString := "The ancient pond. A frog leaps in. Water's sound."
	//testString := "fuck the police. this is the life. i do not know."

	poemFile, err := os.Open("./haiku.json")
	if err != nil {
		fmt.Print("uhoh")
	}
	defer poemFile.Close()

	var poems []string
	byteValue, _ := ioutil.ReadAll(poemFile)
	json.Unmarshal(byteValue, &poems)

	rand.Seed(time.Now().Unix())
	poem := poems[rand.Intn(len(poems))]

	world := board.New(poem)

	generations := len(world.BinaryString)

	fmt.Print(generations)

	fmt.Print("\nInitial world:\n")

	world.String()

	for i := 0; i < generations; i++ {
		world.Progress()
	}

	fmt.Print("Final world:\n")
	world.String()

	fmt.Print(poem)
	fmt.Print('\n')

	world.Draw("basho.png")
}
