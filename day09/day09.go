package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//ReadInput reads line by line from filename and returns a single string
func ReadInput(filename string) string {
	var out string
	fn, err := os.Open(filename) // open "filename" and panics at any error
	if err != nil {
		panic("File Error: " + err.Error())
	}
	defer fn.Close()

	scanner := bufio.NewScanner(fn)
	for scanner.Scan() {
		out += scanner.Text() // reads one line of the file as string and appends to the string

	}
	return out
}

func main() {
	var inGarbage = false
	var score, level, garbageCount int
	var cleanStream []string

	input := strings.Split(ReadInput("input_day09.txt"), "")

	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == "!" {
			i++
		} else if char == "<" {
			inGarbage = true
		} else if inGarbage {
			if char == ">" {
				inGarbage = false
			} else {
				garbageCount++
			}
		} else if char == "{" || char == "}" {
			cleanStream = append(cleanStream, char)
		}
	}
	for _, char := range cleanStream {
		switch char {
		case "{":
			level++
			score += level
		case "}":
			level--
		}
	}
	fmt.Println("Part 1:", score)
	fmt.Println("Part 2:", garbageCount)
}
