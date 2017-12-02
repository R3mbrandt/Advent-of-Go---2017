package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//MinMax returns the min and max value of the given int array
func MinMax(array []int) (int, int) {
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

//ReadInput reads the given puzzle input, splits at tabs and returns an array of integer array
func ReadInput(filename string) [][]int {
	var out [][]int
	fn, err := os.Open(filename) // open "filename" and panics at any error
	if err != nil {
		fmt.Print(err)
		panic("File Error")
	}
	defer fn.Close()

	scanner := bufio.NewScanner(fn)
	for scanner.Scan() {
		lineStr := scanner.Text()                        // reads one line of the file as string
		var t []int                                      // the temporarly inner array
		for _, i := range strings.Split(lineStr, "\t") { // splits at tabs
			n, _ := strconv.Atoi(i) // convert string to int
			t = append(t, n)        // append the number to temp array
		}
		out = append(out, t) // append temp array == one line of file to outer array
	}
	return out
}

func main() {
	var sum int
	input := ReadInput("input_day02.txt")

	// Part One - sum up the min an max values of each line
	for _, o := range input {
		min, max := MinMax(o)
		sum += (max - min)
	}
	fmt.Println("Part 1: ", sum)

	// Part Two - sum up the division of the two only even divisible numbers in each line
	sum = 0
	for _, o := range input {
		for i := 0; i < len(o)-1; i++ { // ??? could the loop part be shorter ???
			for a := i + 1; a < len(o); a++ {
				if o[i]%o[a] == 0 {
					sum += o[i] / o[a]
					break
				} else if o[a]%o[i] == 0 {
					sum += o[a] / o[i]
					break
				}
			}
		}
	}
	fmt.Println("Part 2: ", sum)
}
