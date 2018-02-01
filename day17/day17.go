package main

import (
	"fmt"
)

const input = 363
const iters = 2017

func main() {
	var circBuffer = []int{0}
	var currPos int

	for i := 1; i <= iters; i++ {
		currPos = (currPos+input)%len(circBuffer) + 1
		circBuffer = append(circBuffer[:currPos], append([]int{i}, circBuffer[currPos:]...)...)
	}
	fmt.Println("Part 1:", circBuffer[currPos+1])

	currPos = 0
	size := 1
	var value int

	for i := 1; i <= 50000000; i++ {
		currPos = (currPos+input)%size + 1
		if currPos == 1 {
			value = i
		}
		size++
	}
	fmt.Println("Part 2:", value)
}
