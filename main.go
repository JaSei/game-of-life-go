package main

import (
	"fmt"
	"time"
)

var YMAX int
var XMAX int

func main() {
	world := world{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	XMAX = len(world) - 1
	YMAX = len(world[0]) - 1

	for {
		world.Evaluate()
		world.Print()
		world.NextGeneration()
		fmt.Println()
		time.Sleep(time.Second)
	}
}
