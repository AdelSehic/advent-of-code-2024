package helpers

type Coord struct {
	Y int
	X int
}

func (c *Coord) Right() *Coord {
	return &Coord{c.Y, c.X + 1}
}

func (c *Coord) Left() *Coord {
	return &Coord{c.Y, c.X - 1}
}

func (c *Coord) Up() *Coord {
	return &Coord{c.Y - 1, c.X}
}

func (c *Coord) Down() *Coord {
	return &Coord{c.Y + 1, c.X}
}

func (c *Coord) TopLeft() *Coord {
	return c.Up().Left()
}

func (c *Coord) TopRight() *Coord {
	return c.Up().Right()
}

func (c *Coord) BottomLeft() *Coord {
	return c.Down().Left()
}

func (c *Coord) BottomRight() *Coord {
	return c.Down().Right()
}

func (c *Coord) neighbors() []*Coord {
	return []*Coord{
		c.TopLeft(), c.Up(), c.TopRight(),
		c.Left(), c.Right(),
		c.BottomLeft(), c.Down(), c.BottomRight(),
	}
}
