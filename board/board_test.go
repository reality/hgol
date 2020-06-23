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

	if len(board.binaryStringBinaryString) != expectedBinaryStringLength {
		t.Fatalf("Wrong length of binary string. got=%d, expected=%d", len(board.binaryStringBinaryString), expectedBinaryStringLength)
	}

	if len(board.w) != expectedY {
		t.Fatalf("Wrong x size of board. got=%d, expected=%d", len(board.w), expectedY)
	}

	if len(board.w[0]) != expectedX {
		t.Fatalf("Wrong x size of board. got=%d, expected=%d", len(board.w[0]), expectedX)
	}

	neighbourTests := []struct {
		y                  int
		x                  int
		expectedNeighbours []rune
	}{
		{0, 0, []rune{'0', '0', '0', '0', '1', '1', '0', '1'}},
		{10, 15, []rune{'0', '1', '1', '0', '0', '0', '0', '0'}},
		{8, 0, []rune{'0', '0', '1', '0', '1', '0', '0', '1'}},
	}

	for _, tt := range neighbourTests {
		neighbours := board.getNeighbours(tt.y, tt.x)
		if string(tt.expectedNeighbours) != string(neighbours) {
			t.Fatalf("Neighbours for y=%d,x=%d were not as expected.\nexpected=%s\ngot=     %s", tt.y, tt.x,
				string(tt.expectedNeighbours),
				string(neighbours))
		}
	}

	fmt.Printf("%c", board.w[10][15])

	board.Progress()

	progressTests := []struct {
		y             int
		x             int
		expectedValue rune
	}{
		{0, 0, '1'},   // remain alive
		{10, 15, '0'}, // remain dead
		{8, 0, '1'},   // become alive
	}

	for _, tt := range progressTests {
		if board.w[tt.y][tt.x] != tt.expectedValue {
			t.Fatalf("Progression for y=%d,x=%d was not as expected. expected=%c, got=%c.", tt.y, tt.x,
				tt.expectedValue, board.w[tt.y][tt.x])
		}
	}
}
