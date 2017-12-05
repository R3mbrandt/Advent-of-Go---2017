package main

import (
	"fmt"
	"math"
)

const (
	right = iota
	up
	left
	down
	rightUp
	leftDown
	leftUp
	rightDown
)

const target = 347991

var dimension = int(math.Sqrt(target)) + 2
var dirs = [...]int{right, up, left, down, rightUp, leftDown, leftUp, rightDown}

// move moves "x,y" values in direction "dir" and returns the new position
func move(x, y, dir int) (int, int) { // Quite unhappy with this part!
	switch dir {
	case right:
		x++
	case up:
		y++
	case left:
		x--
	case down:
		y--
	case rightUp:
		x++
		y++
	case leftDown:
		x--
		y--
	case leftUp:
		x--
		y++
	case rightDown:
		x++
		y--
	}

	return x, y
}

// sum the value of all neighboors in grid at position x,y
func sumNeighboors(x, y int, grid [][]int) int {
	sum := 0
	for dir := range dirs {
		x1, y1 := move(x, y, dir)
		sum += grid[x1][y1]
	}
	return sum
}

func main() {
	grid := make([][]int, dimension) // Quite unhappy with this part, too!
	for i := 0; i < dimension; i++ {
		grid[i] = make([]int, dimension)
	}
	start := dimension / 2
	x, y := start, start
	dir := dirs[0]

	for i := 0; i < target-1; i++ {
		grid[x][y] = i + 1
		nextDir := dirs[(dir+1)%4]
		x1, y1 := move(x, y, nextDir)
		if grid[x1][y1] == 0 {
			dir = nextDir
		}
		x, y = move(x, y, dir)
	}

	dist := math.Abs(float64(x-start)) + math.Abs(float64(y-start))
	fmt.Println("Part 1: ", dist)

	// Part 2

	grid = make([][]int, dimension)
	for i := 0; i < dimension; i++ {
		grid[i] = make([]int, dimension)
	}
	x, y = start, start
	dir = dirs[0]
	grid[x][y] = 1

	for {
		nextDir := dirs[(dir+1)%4]
		x1, y1 := move(x, y, nextDir)
		if grid[x1][y1] == 0 {
			dir = nextDir
		}
		x, y = move(x, y, dir)
		grid[x][y] = sumNeighboors(x, y, grid)
		if grid[x][y] > target {
			fmt.Println("Part 2: ", grid[x][y])
			break
		}
	}
}
