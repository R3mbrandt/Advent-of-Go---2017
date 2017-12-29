package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Layer struct {
	rnge   int
	dir    int
	curPos int
}

//ReadInput reads line by line from filename and returns a map of layers
func ReadInput(filename string) map[int]Layer {
	out := make(map[int]Layer)

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
		out[lyer] = Layer{rnge: rnge, dir: 1, curPos: 0}
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

func makeOneStep(input map[int]Layer) {
	for key, value := range input {
		direc := value.dir
		if value.curPos+value.dir >= value.rnge || value.curPos+value.dir < 0 {
			direc *= -1
		}
		input[key] = Layer{rnge: value.rnge, dir: direc, curPos: value.curPos + direc}
	}
}

func main() {
	var sum int
	input := ReadInput("input_day13.txt")
	keys := make([]int, 0, len(input))

	for k := range input {
		keys = append(keys, k)
	}
	maxK := Max(keys)
	//fmt.Println(input, Max(keys))

	for i := 0; i <= maxK; i++ {
		if v, ok := input[i]; ok {
			if v.curPos == 0 {
				sum += i * v.rnge
			}
		}
		makeOneStep(input)
		//	fmt.Println(input)
	}
	fmt.Println("Part 1:", sum)

	var delay int
	var caught bool
	for {
		input = ReadInput("input_day13.txt")
		for step := 0; step < delay; step++ {
			makeOneStep(input)
		}
		caught = false
		for i := 0; i <= maxK; i++ {
			if v, ok := input[i]; ok {
				if v.curPos == 0 {
					caught = true
					break
				}
			}
			makeOneStep(input)
		}
		if !caught {
			break
		}
		delay++
	}
	fmt.Println("Part 2:", delay)
}
