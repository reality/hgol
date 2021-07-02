package world

import (
	"fmt"
)

type World struct {
	w [][]rune // the w-orld
}

func New(sizeX, sizeY int) *World {

	// Create the board with the given size
	b := &World{}
	b.w = make([][]rune, sizeX)
	for i := range b.w {
		b.w[i] = make([]rune, sizeY)
	}

	// Now we need to initialise the world with tiles
	// So I think the general idea is that we want to create sort of 'veins' of different habitats. I guess we can look up ways to do that.
	// for now we will make it into 0

	for x, col := range b.w {
		for y, _ := range col {
			b.w[x][y] = '0'
		}
	}

	return b
}

func (b *World) Progress() {
	newWorld := make([][]rune, len(b.w))
	for i := range b.w {
		newWorld[i] = make([]rune, len(b.w[i]))
		copy(newWorld[i], b.w[i])
	}

	b.w = newWorld
}

// Get the eight neighbours of a cell position.
// Essentially we either look up the value of the cell, or if it's outside the map, we assume it's dead
// TODO this could potentially be simplified with a getCell that has a default value
func (b *World) getNeighbours(y int, x int) []rune {
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

// currently the wrong way around
func (b *World) String() {
	for _, row := range b.w {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Printf("\n")
	}
}
