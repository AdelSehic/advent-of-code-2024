package main

import "github.com/AdelSehic/advent-of-code-2024/helpers"

const (
	FACING_UP    = "up"
	FACING_DOWN  = "down"
	FACING_LEFT  = "LEFT"
	FACING_RIGHT = "RIGHT"
)

type Guard struct {
	MoveFunc    func(*helpers.Coord) *helpers.Coord
	Facing      string
	Position    *helpers.Coord
	OriginalPos *helpers.Coord
}

func NewGuard(location *helpers.Coord) *Guard {
	return &Guard{
		MoveFunc:    (*helpers.Coord).Up,
		Position:    location,
		Facing:      FACING_UP,
		OriginalPos: location,
	}
}

func (g *Guard) Copy() *Guard {
	return &Guard{
		MoveFunc:    g.MoveFunc,
		Position:    g.Position,
		Facing:      g.Facing,
		OriginalPos: g.OriginalPos,
	}
}

func (g *Guard) Reset() {
	g.Facing = FACING_UP
	g.MoveFunc = (*helpers.Coord).Up
	g.Position = g.OriginalPos
}

func (g *Guard) NextField() *helpers.Coord {
	return g.MoveFunc(g.Position)
}

func (g *Guard) InFront(f *helpers.Field) byte {
	return f.GetLetter(g.NextField())
}

func (g *Guard) Move() {
	g.Position = g.NextField()
}

func (g *Guard) Rotate() {
	switch g.Facing {
	case FACING_UP:
		g.MoveFunc = (*helpers.Coord).Right
		g.Facing = FACING_RIGHT
	case FACING_RIGHT:
		g.MoveFunc = (*helpers.Coord).Down
		g.Facing = FACING_DOWN
	case FACING_DOWN:
		g.MoveFunc = (*helpers.Coord).Left
		g.Facing = FACING_LEFT
	case FACING_LEFT:
		g.MoveFunc = (*helpers.Coord).Up
		g.Facing = FACING_UP
	}
}
