package main

import (
	"bufio"
	"fmt"
	//"math"
	"os"
	"strconv"
	"strings"
)

//ReadInput reads line by line from filename and returns a map[]int of []int
func ReadInput(filename string) map[int][]int {
	out := make(map[int][]int)
	fn, err := os.Open(filename) // open "filename" and panics at any error
	if err != nil {
		panic("File Error: " + err.Error())
	}
	defer fn.Close()

	scanner := bufio.NewScanner(fn)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " <-> ")
		prog, _ := strconv.Atoi(tmp[0])
		values := strings.Split(tmp[1], ",")
		for _, v := range values {
			t, _ := strconv.Atoi(strings.TrimSpace(v))
			out[prog] = append(out[prog], t)
		}
	}
	return out
}

func main() {
	input := ReadInput("input_day12.txt")
	queue := []int{0}
	group := make(map[int]bool)

	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		if _, ok := group[elem]; !ok {
			group[elem] = true
			queue = append(queue, input[elem]...)
		}
	}
	fmt.Println("Part 1:", len(group))

	var nGroups int
	for keys := range input {
		nGroups++
		queue = []int{keys}
		group = make(map[int]bool)
		for len(queue) > 0 {
			elem := queue[0]
			queue = queue[1:]
			if _, ok := group[elem]; !ok {
				group[elem] = true
				queue = append(queue, input[elem]...)
				delete(input, elem)
			}
		}
	}
	fmt.Println("Part 2:", nGroups)
}
