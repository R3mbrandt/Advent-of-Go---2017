package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//ReadInput reads line by line from filename and returns a single string
func ReadInput(filename string) (out string) {
	fn, err := os.Open(filename) // open "filename" and panics at any error
	if err != nil {
		panic("File Error: " + err.Error())
	}
	defer fn.Close()

	scanner := bufio.NewScanner(fn)
	for scanner.Scan() {
		out += scanner.Text() // reads one line of the file as string and appends to the string

	}
	return
}

func main() {
	var inGarbage = false
	var esc = false
	var score, level, garbageCount int

	input := strings.Split(ReadInput("input_day09.txt"), "")

	for _, char := range input {
		if inGarbage {
			if esc {
				esc = false
			} else if char == ">" {
				inGarbage = false
			} else if char == "!" {
				esc = true
			} else {
				garbageCount++
			}
		} else if char == "{" {
			level++
			score += level
		} else if char == "}" {
			level--
		} else if char == "<" {
			inGarbage = true
		}
	}
	fmt.Println("Part 1:", score)
	fmt.Println("Part 2:", garbageCount)
}
