package helpers

type Coord struct {
	Y int
	X int
}

func (c *Coord) right() *Coord {
	return &Coord{c.Y, c.X + 1}
}

func (c *Coord) left() *Coord {
	return &Coord{c.Y, c.X - 1}
}

func (c *Coord) up() *Coord {
	return &Coord{c.Y - 1, c.X}
}

func (c *Coord) down() *Coord {
	return &Coord{c.Y + 1, c.X}
}

// Get top-left edge (0, 0)
func (c *Coord) topLeft() *Coord {
	return c.up().left()
}

// Get top-right edge
func (c *Coord) topRight() *Coord {
	return c.up().right()
}

// Get bottom-left edge
func (c *Coord) bottomLeft() *Coord {
	return c.down().left()
}

// Get bottom-right edge
func (c *Coord) bottomRight() *Coord {
	return c.down().right()
}

// Get all neighbors
func (c *Coord) neighbors() []*Coord {
	return []*Coord{
		c.topLeft(), c.up(), c.topRight(),
		c.left(), c.right(),
		c.bottomLeft(), c.down(), c.bottomRight(),
	}
}
