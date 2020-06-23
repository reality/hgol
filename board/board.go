package board

import "fmt"

const (
	MAX_X = 50 // 20 * 8, so we should get 20 characters on each line
)

type Board struct {
	w            [][]rune // the w-orld
	inputString  string
	binaryString string
}

func New(input string) *Board {
	b := &Board{
		inputString: input,
	}

	// Build the binary string
	for _, c := range b.inputString {
		b.binaryString = fmt.Sprintf("%s%08b", b.binaryString, c)
	}

	// Create the board
	boardX := MAX_X
	boardY := int(len(b.binaryString)/boardX) + 1

	b.w = make([][]rune, boardY)
	for i := range b.w {
		b.w[i] = make([]rune, boardX)
	}

	// Initialise the board with the binarised string
	for pos, char := range b.binaryString {
		charY, charX := pos/boardX, pos%boardX
		b.w[charY][charX] = char
	}
	// TODO finish off the final line with the \x00 content

	return b
}

func (b *Board) progress() {}

// Get the eight neighbours of a cell position.
// Essentially we either look up the value of the cell, or if it's outside the map, we assume it's dead
// TODO this could potentially be simplified with a getCell that has a default value
func (b *Board) getNeighbours(x int, y int) []rune {
	var neighbours []rune

	// First we iterate the three cells above our target cell, then the three below

	above := y - 1
	if above >= 0 && above < len(b.w) {
		for tX := x - 1; tX <= x+1; tX++ {
			if tX >= 0 && tX < len(b.w[above]) {
				neighbours = append(neighbours, b.w[above][tX])
			} else {
				neighbours = append(neighbours, '0')
			}
		}
	} else {
		neighbours = append(neighbours, []rune{'0', '0', '0'}...)
	}

	below := y + 1
	if below >= 0 && below < len(b.w) {
		for tX := x - 1; tX <= x+1; tX++ {
			if tX >= 0 && tX < len(b.w[below]) {
				neighbours = append(neighbours, b.w[below][tX])
			} else {
				neighbours = append(neighbours, '0')
			}
		}
	} else {
		neighbours = append(neighbours, []rune{'0', '0', '0'}...)
	}

	// Now we'll just get x+1 and x-1 manually
	previous := x - 1
	if previous > 0 && previous < len(b.w[y]) {
		neighbours = append(neighbours, b.w[y][previous])
	} else {
		neighbours = append(neighbours, '0')
	}

	following := x + 1
	if following > 0 && following < len(b.w[y]) {
		neighbours = append(neighbours, b.w[y][following])
	} else {
		neighbours = append(neighbours, '0')
	}

	return neighbours
}

func (b *Board) draw() {}

func (b *Board) String() {
	for _, row := range b.w {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Printf("\n")
	}
}
