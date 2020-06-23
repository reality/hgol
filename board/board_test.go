package board

import (
	"fmt"
	"testing"
)

func TestBoard(t *testing.T) {
	testString := "Hello this is dog, I love to run and to love, I get left at home"
	expectedBinaryStringLength := len(testString) * 8
	expectedX := MAX_X
	expectedY := int(expectedBinaryStringLength/MAX_X) + 1

	fmt.Printf("%s\n", testString)

	board := New(testString)

	board.String()

	if len(board.binaryString) != expectedBinaryStringLength {
		t.Fatalf("Wrong length of binary string. got=%d, expected=%d", len(board.binaryString), expectedBinaryStringLength)
	}

	if len(board.w) != expectedY {
		t.Fatalf("Wrong x size of board. got=%d, expected=%d", len(board.w), expectedY)
	}

	if len(board.w[0]) != expectedX {
		t.Fatalf("Wrong x size of board. got=%d, expected=%d", len(board.w[0]), expectedX)
	}

	neighbourTests := []struct {
		x                  int
		y                  int
		expectedNeighbours []rune
	}{
		{0, 0, []rune{'0', '0', '0', '0', '1', '1', '0', '1'}},
	}

	for _, tt := range neighbourTests {
		neighbours := board.getNeighbours(tt.x, tt.y)
		if string(tt.expectedNeighbours) != string(neighbours) {
			t.Fatalf("Neighbours for y=%d,x=%d were not as expected.\nexpected=%s\ngot=     %s", tt.y, tt.x,
				string(tt.expectedNeighbours),
				string(neighbours))
		}
	}

	board.progress()

	board.String()
}
