package main

type coord struct {
	y int
	x int
}

func (c *coord) right() *coord {
	return &coord{c.y, c.x + 1}
}

func (c *coord) left() *coord {
	return &coord{c.y, c.x - 1}
}

func (c *coord) up() *coord {
	return &coord{c.y - 1, c.x}
}

func (c *coord) down() *coord {
	return &coord{c.y + 1, c.x}
}

// Get top-left edge (0, 0)
func (c *coord) topLeft() *coord {
	return c.up().left()
}

// Get top-right edge
func (c *coord) topRight() *coord {
	return c.up().right()
}

// Get bottom-left edge
func (c *coord) bottomLeft() *coord {
	return c.down().left()
}

// Get bottom-right edge
func (c *coord) bottomRight() *coord {
	return c.down().right()
}

// Get all neighbors
func (c *coord) neighbors() []*coord {
	return []*coord{
		c.topLeft(), c.up(), c.topRight(),
		c.left(), c.right(),
		c.bottomLeft(), c.down(), c.bottomRight(),
	}
}
