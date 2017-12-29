package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//ReadInput reads line by line from filename and returns a map of ints
func ReadInput(filename string) map[int]int {
	out := make(map[int]int)

	fn, err := os.Open(filename) // open "filename" and panics at any error
	if err != nil {
		fmt.Print(err)
		panic("File Error")
	}
	defer fn.Close()

	scanner := bufio.NewScanner(fn)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), ":")
		lyer, _ := strconv.Atoi(tmp[0])
		rnge, _ := strconv.Atoi(strings.TrimSpace(tmp[1]))
		out[lyer] = rnge
	}
	return out
}

func Max(array []int) int {
	var max = array[0]

	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func main() {
	var sum int
	input := ReadInput("input_day13.txt")
	keys := make([]int, 0, len(input))

	for k := range input {
		keys = append(keys, k)
	}
	maxK := Max(keys)

	for i := 0; i <= maxK; i++ {
		if v, ok := input[i]; ok {
			if i%((v-1)*2) == 0 {
				sum += i * v
			}
		}
	}
	fmt.Println("Part 1:", sum)

	var delay int
	var caught bool
	for {
		caught = false
		for i := 0; i <= maxK; i++ {
			if v, ok := input[i]; ok {
				if (delay+i)%((v-1)*2) == 0 {
					caught = true
					break
				}
			}
		}
		if !caught {
			break
		}
		delay++
	}
	fmt.Println("Part 2:", delay)
}
