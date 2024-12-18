package main

import (
	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

type PathTracer struct {
	Path  []*helpers.Coord
	Iter  *helpers.FieldIterator
	Steps int
	Turns int
}

func (p *PathTracer) Pos() *helpers.Coord {
	return p.Iter.Position
}

func NewTracer(place *helpers.Coord, places ...*helpers.Coord) *PathTracer {
	tracer := &PathTracer{
		Path:  make([]*helpers.Coord, 0, len(places)),
		Iter:  helpers.NewFieldIterator(place),
		Steps: 0,
		Turns: 1,
	}
	tracer.Path = append(tracer.Path, places...)
	return tracer.Rotate().Rotate()
}

func (path *PathTracer) Copy() *PathTracer {
	newPath := &PathTracer{
		Path:  make([]*helpers.Coord, 0, len(path.Path)),
		Steps: path.Steps,
		Turns: path.Turns,
	}
	newPath.Iter = path.Iter.Copy()
	newPath.Path = append(newPath.Path, path.Path...)

	return newPath
}

func (pt *PathTracer) RotateOther() *PathTracer {
	return pt.rotateHelper((*helpers.FieldIterator).RotateOther)
}

func (pt *PathTracer) Rotate() *PathTracer {
	return pt.rotateHelper((*helpers.FieldIterator).Rotate)
}

func (pt *PathTracer) rotateHelper(fn func(*helpers.FieldIterator)) *PathTracer {
	newPath := pt.Copy()
	fn(newPath.Iter)
	newPath.Turns++
	newPath.Path = append(newPath.Path, newPath.Iter.Position)
	return newPath
}

func (pt *PathTracer) Move() {
	pt.Iter.Move()
	pt.Path = append(pt.Path, pt.Iter.Position)
	pt.Steps++
}

func (pt *PathTracer) Solution() int {
	sol := pt.Steps + (pt.Turns * 1000)
	return sol
}
