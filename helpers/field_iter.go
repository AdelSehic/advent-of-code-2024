package helpers

const (
	FACING_UP    = "up"
	FACING_DOWN  = "down"
	FACING_LEFT  = "LEFT"
	FACING_RIGHT = "RIGHT"
)

type FieldIterator struct {
	MoveFunc    func(*Coord) *Coord
	Facing      string
	Position    *Coord
	OriginalPos *Coord
}

func NewFieldIterator(location *Coord) *FieldIterator {
	return &FieldIterator{
		MoveFunc:    (*Coord).Up,
		Position:    location,
		Facing:      FACING_UP,
		OriginalPos: location,
	}
}

func (g *FieldIterator) Copy() *FieldIterator {
	return &FieldIterator{
		MoveFunc:    g.MoveFunc,
		Position:    g.Position,
		Facing:      g.Facing,
		OriginalPos: g.OriginalPos,
	}
}

func (g *FieldIterator) Reset() {
	g.Facing = FACING_UP
	g.MoveFunc = (*Coord).Up
	g.Position = g.OriginalPos
}

func (g *FieldIterator) NextField() *Coord {
	return g.MoveFunc(g.Position)
}

func (g *FieldIterator) InFront(f *Field) byte {
	return f.GetLetter(g.NextField())
}

func (g *FieldIterator) Move() {
	g.Position = g.NextField()
}

func (g *FieldIterator) Rotate() {
	switch g.Facing {
	case FACING_UP:
		g.MoveFunc = (*Coord).Right
		g.Facing = FACING_RIGHT
	case FACING_RIGHT:
		g.MoveFunc = (*Coord).Down
		g.Facing = FACING_DOWN
	case FACING_DOWN:
		g.MoveFunc = (*Coord).Left
		g.Facing = FACING_LEFT
	case FACING_LEFT:
		g.MoveFunc = (*Coord).Up
		g.Facing = FACING_UP
	}
}

func NewMoveFunc(x, y int) func(*Coord) *Coord {
	return func(crd *Coord) *Coord {
		return &Coord{
			X: crd.X + x,
			Y: crd.Y + y,
		}
	}
}
