package world

type Object struct {
	o       [][]rune
	trimmed bool // Has the object been trimmed using the Trim function?
}

func CreateObject(sizeX, sizeY int) *Object {
	// Create the board with the given size
	b := &Object{}
	b.o = make([][]rune, sizeX)
	for i := range b.o {
		b.o[i] = make([]rune, sizeY)
	}

	for x, col := range b.o {
		for y, _ := range col {
			b.o[x][y] = '0'
		}
	}

	return b
}

func (o *Object) Draw(world World) {

}

// Here we cut off any unused space. So we basically just have to find the
// left/right/bottom/top boundary and cut things off past that point. TODO
func (o *Object) Trim() {
	o.trimmed = true
}
