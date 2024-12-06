package main

import "github.com/AdelSehic/advent-of-code-2024/helpers"

const (
	FACING_UP    = "up"
	FACING_DOWN  = "down"
	FACING_LEFT  = "LEFT"
	FACING_RIGHT = "RIGHT"
)

type Guard struct {
	MoveFunc func(*helpers.Coord) *helpers.Coord
	Facing   string
	Positing *helpers.Coord
}

func (g *Guard) InFront(f *helpers.Field) byte {
	return f.GetLetter(g.MoveFunc(g.Positing))
}

func (g *Guard) Move() {
	g.Positing = g.MoveFunc(g.Positing)
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
