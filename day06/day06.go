package main

import (
	"fmt"
	"strings"
)

//FindPosMax find the biggest value in array and returns the position (of first occurance)
func FindPosMax(array []int) int {
	var max = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	for pos, value := range array {
		if value == max {
			return pos
		}
	}
	return 0
}

//IntToStr converts an array of ints into a string
func IntToStr(input []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(input), " ", "", -1), "[]")
}

func main() {
	input := []int{4, 1, 15, 12, 0, 9, 9, 5, 5, 8, 7, 3, 14, 5, 12, 3}
	hist := make(map[string]int)
	counter := 0

	for {
		counter++
		bigPos := FindPosMax(input)
		bigVal := input[bigPos]
		input[bigPos] = 0
		for i := bigVal; i > 0; i-- {
			bigPos = (bigPos + 1) % len(input)
			input[bigPos]++
		}
		key := IntToStr(input)
		if _, ok := hist[key]; ok {
			fmt.Println("Part 2: ", counter-hist[key])
			break
		} else {
			hist[key] = counter
		}
	}
	fmt.Println("Part 1: ", counter)
}
