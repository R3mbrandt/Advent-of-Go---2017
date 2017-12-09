package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//ReadInput reads line by line from filename and returns an string array
func ReadInput(filename string) []string {
	var out []string
	fn, err := os.Open(filename) // open "filename" and panics at any error
	if err != nil {
		fmt.Print(err)
		panic("File Error")
	}
	defer fn.Close()

	scanner := bufio.NewScanner(fn)
	for scanner.Scan() {
		out = append(out, scanner.Text()) // reads one line of the file as string and appends to the string array

	}
	return out
}

func main() {
	parents := make(map[string]bool)
	childs := make(map[string]bool)

	input := ReadInput("input_day07.txt")
	re := regexp.MustCompile(`(.*)\s?\((\d+)\)\s?(?:->)?\s?(.*)`)

	for _, line := range input {
		result := re.FindStringSubmatch(line)
		if result[1] != "" {
			parents[strings.TrimSpace(result[1])] = true
		}
		if result[3] != "" {
			for _, k := range strings.Split(result[3], ", ") {
				childs[k] = true
			}
		}
	}
	for key := range childs {
		delete(parents, key)
	}

	for key := range parents {
		fmt.Println("Part 1:", key)
	}
}
