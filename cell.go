package main

type cell byte

func (c cell) IsAlive() bool {
	return c&1 == 1
}

func (c *cell) Revival() {
	*c |= 1
}

func (c *cell) Die() {
	*c &= 0
}

func (c cell) GetNeighbors() byte {
	return byte(c >> 1)
}

func (c *cell) SetNeighbors(neighbors byte) {
	*c = (*c << 7) >> 7
	*c |= cell(neighbors << 1)
}
