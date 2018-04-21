package main

import (
	"fmt"
)

type world struct {
	endx     int
	endy     int
	latest   [][]cell
	previous [][]cell
}

func New(endx, endy int) world {
	return world{endx: endx, endy: endy, latest: newS(endx, endy), previous: newS(endx, endy)}
}

func newS(endx, endy int) [][]cell {
	s := make([][]cell, endx+1)
	for i := range s {
		s[i] = make([]cell, endy+1)
	}

	return s
}

func (w world) Print() {
	for _, x := range w.latest {
		for _, cell := range x {
			if cell.IsAlive() {
				fmt.Print("#")
			} else {
				fmt.Print("_")
			}
		}

		fmt.Print(" | ")

		for _, cell := range x {
			fmt.Print(cell.GetNeighbors())
		}

		fmt.Println()
	}
}

func (w *world) Evaluate() {
	for x, row := range (*w).latest {
		for y, _ := range row {
			(*w).latest[x][y].SetNeighbors(w.Neighbors(x, y))
		}
	}
}

func (w *world) NextGeneration() {
	for i, row := range w.latest {
		for j, val := range row {
			w.previous[i][j] = val
		}
	}

	for x, row := range (*w).latest {
		for y, cell := range row {
			n := cell.GetNeighbors()
			if n < 2 || n > 3 {
				(*w).latest[x][y].Die()
			} else if n == 3 {
				(*w).latest[x][y].Revival()
			}
		}
	}
}

func (w world) Neighbors(x, y int) (neighbors byte) {
	if x > 0 && w.latest[x-1][y].IsAlive() {
		neighbors++
	}
	if x > 0 && y > 0 && w.latest[x-1][y-1].IsAlive() {
		neighbors++
	}
	if x > 0 && y < w.endy && w.latest[x-1][y+1].IsAlive() {
		neighbors++
	}
	if x < w.endx && w.latest[x+1][y].IsAlive() {
		neighbors++
	}
	if x < w.endx && y > 0 && w.latest[x+1][y-1].IsAlive() {
		neighbors++
	}
	if x < w.endx && y < w.endy && w.latest[x+1][y+1].IsAlive() {
		neighbors++
	}
	if y > 0 && w.latest[x][y-1].IsAlive() {
		neighbors++
	}
	if y < w.endy && w.latest[x][y+1].IsAlive() {
		neighbors++
	}

	return
}

func (w world) IsStable() bool {
	for x, row := range w.latest {
		for y, _ := range row {
			if w.latest[x][y].IsAlive() != w.previous[x][y].IsAlive() {
				return false
			}
		}
	}

	return true
}
