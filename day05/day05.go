package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//ReadInput reads the given puzzle input, line by line and returns an array of integer
func ReadInput(filename string) []int {
	var out []int
	fn, err := os.Open(filename) // open "filename" and panics at any error
	if err != nil {
		fmt.Print(err)
		panic("File Error")
	}
	defer fn.Close()

	scanner := bufio.NewScanner(fn)
	for scanner.Scan() {
		t, _ := strconv.Atoi(scanner.Text())
		out = append(out, t)
	}
	return out
}

func main() {
	var input, input2 []int

	input = ReadInput("input_day05.txt")
	input2 = append(input2, input...)

	currPos := 0
	counter := 0
	oldPos := 0

	for {
		counter++

		if currPos+input[currPos] > len(input)-1 {
			break
		} else {
			oldPos = currPos
			currPos += input[currPos]
			input[oldPos]++
		}
	}
	fmt.Println(counter)

	currPos = 0
	counter = 0
	oldPos = 0

	for {
		counter++

		if currPos+input2[currPos] > len(input2)-1 {
			break
		} else {
			oldPos = currPos
			currPos += input2[currPos]
			if input2[oldPos] >= 3 {
				input2[oldPos]--
			} else {
				input2[oldPos]++
			}
		}
	}
	fmt.Println(counter)
}
