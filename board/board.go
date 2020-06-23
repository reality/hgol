package board

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"

	"github.com/muesli/gamut"
	//"github.com/muesli/gamut/palette"
	//"github.com/muesli/gamut/theme"
)

const (
	MAX_X  = 24
	BLOWUP = 25
)

type Board struct {
	w            [][]rune // the w-orld
	inputString  string
	BinaryString string
}

func New(input string) *Board {
	b := &Board{
		inputString: input,
	}

	// Build the binary string
	for _, c := range b.inputString {
		b.BinaryString = fmt.Sprintf("%s%08b", b.BinaryString, c)
	}

	// Create the board
	boardX := MAX_X
	boardY := int(len(b.BinaryString)/boardX) + 1

	b.w = make([][]rune, boardY)
	for i := range b.w {
		b.w[i] = make([]rune, boardX)
	}

	// Initialise the board with the binarised string
	for pos, char := range b.BinaryString {
		charY, charX := pos/boardX, pos%boardX
		b.w[charY][charX] = char
	}

	for pos, char := range b.w[boardY-1] {
		if char != '0' && char != '1' {
			b.w[boardY-1][pos] = '0'
		}
	}

	return b
}

func (b *Board) Progress() {
	newWorld := make([][]rune, len(b.w))
	for i := range b.w {
		newWorld[i] = make([]rune, len(b.w[i]))
		copy(newWorld[i], b.w[i])
	}

	for y, row := range b.w {
		for x, cell := range row {
			neighbours := b.getNeighbours(y, x)
			aliveNeighbours := 0

			for _, nVal := range neighbours {
				if nVal == '1' {
					aliveNeighbours++
				}
			}

			//fmt.Printf("y=%d,x=%d. alive neighbours=%d\n", y, x, aliveNeighbours)

			// a dead cell with three living neighbours becomes alive, a living cell with two or three living neighbours can remain alive, but otherwise, everything must die
			if cell == '0' && aliveNeighbours == 3 {
				cell = '1'
			} else if !(cell == '1' && (aliveNeighbours == 2 || aliveNeighbours == 3)) {
				cell = '0'
			}

			newWorld[y][x] = cell
		}
	}

	b.w = newWorld
}

// Get the eight neighbours of a cell position.
// Essentially we either look up the value of the cell, or if it's outside the map, we assume it's dead
// TODO this could potentially be simplified with a getCell that has a default value
func (b *Board) getNeighbours(y int, x int) []rune {
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

// Draw the image. Bit screwed up here because it's er, x,y rather than y,x
func (b *Board) Draw(fileName string) {
	g := gg.NewContext(len(b.w[0])*BLOWUP, len(b.w)*BLOWUP)

	g.DrawRectangle(0, 0, float64(len(b.w[0])*BLOWUP), float64(len(b.w)*BLOWUP))
	g.SetRGB(0, 0, 0)
	g.Fill()

	rCol := int(b.inputString[0] - '0')
	gCol := int(b.inputString[1] - '0')
	bCol := int(b.inputString[2] - '0')

	livingCellColor := color.RGBA{uint8(rCol), uint8(gCol), uint8(bCol), 1}

	g.DrawRectangle(0, 0, float64(len(b.w[0])*BLOWUP), float64(len(b.w)*BLOWUP))

	shades := gamut.Shades(livingCellColor, 5)
	bgR, bgG, bgB, _ := shades[3].RGBA()
	g.SetRGB(float64(bgR), float64(bgG), float64(bgB))

	g.Fill()

	for y, row := range b.w {
		for x, cell := range row {
			if cell == '1' {
				g.DrawRectangle(float64(x*BLOWUP), float64(y*BLOWUP), 20, 20)
				g.SetRGB(float64(rCol), float64(gCol), float64(bCol))
				g.Fill()
				fmt.Printf("drawing white at %f %f\n", float64(y*BLOWUP), float64(x*BLOWUP))
			}
		}
	}

	g.SavePNG(fileName)
}

func (b *Board) String() {
	for _, row := range b.w {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Printf("\n")
	}
}
