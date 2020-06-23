package board

import "fmt"

const (
	MAX_X = 160 // 20 * 8, so we should get 20 characters on each line
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

func (b *Board) getNeighbours() {}

func (b *Board) draw() {}

func (b *Board) String() {
	for _, row := range b.w {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Printf("\n")
	}
}
