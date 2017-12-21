package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

//ReadInput reads line by line from filename and returns an string array
func ReadInput(filename string) []string {
	var out []string
	fn, err := os.Open(filename) // open "filename" and panics at any error
	if err != nil {
		panic("File Error: " + err.Error())
	}
	defer fn.Close()

	scanner := bufio.NewScanner(fn)
	for scanner.Scan() {
		out = append(out, scanner.Text()) // reads one line of the file as string and appends to the string array

	}
	return out
}

func Max(array []float64) float64 {
	var max = array[0]

	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func main() {
	var x, y int
	var max float64
	input := strings.Split(ReadInput("input_day11.txt")[0], ",")

	for _, e := range input {
		switch e {
		case "n":
			y--
		case "s":
			y++
		case "ne":
			x++
		case "se":
			y++
			x++
		case "nw":
			y--
			x--
		case "sw":
			x--
		}
		steps := []float64{math.Abs(float64(x - y)), math.Abs(float64(x)), math.Abs(float64(y)), max}
		max = Max(steps)

	}
	steps := []float64{math.Abs(float64(x - y)), math.Abs(float64(x)), math.Abs(float64(y))}
	fmt.Println("Part 1:", Max(steps))
	fmt.Println("Part 2:", max)
}
