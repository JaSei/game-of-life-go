package main

import (
	"fmt"
)

type world [XMAX + 1][YMAX + 1]cell

func (w world) Print() {
	for _, x := range w {
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
	for x, row := range *w {
		for y, _ := range row {
			(*w)[x][y].SetNeighbors(w.Neighbors(x, y))
		}
	}
}

func (w *world) NextGeneration() {
	for x, row := range *w {
		for y, cell := range row {
			n := cell.GetNeighbors()
			if n < 2 || n > 3 {
				(*w)[x][y].Die()
			} else if n == 3 {
				(*w)[x][y].Revival()
			}
		}
	}
}

func (w world) Neighbors(x, y int) (neighbors byte) {
	if x > 0 && w[x-1][y].IsAlive() {
		neighbors++
	}
	if x > 0 && y > 0 && w[x-1][y-1].IsAlive() {
		neighbors++
	}
	if x > 0 && y < YMAX && w[x-1][y+1].IsAlive() {
		neighbors++
	}
	if x < XMAX && w[x+1][y].IsAlive() {
		neighbors++
	}
	if x < XMAX && y > 0 && w[x+1][y-1].IsAlive() {
		neighbors++
	}
	if x < XMAX && y < YMAX && w[x+1][y+1].IsAlive() {
		neighbors++
	}
	if y > 0 && w[x][y-1].IsAlive() {
		neighbors++
	}
	if y < YMAX && w[x][y+1].IsAlive() {
		neighbors++
	}

	return
}
