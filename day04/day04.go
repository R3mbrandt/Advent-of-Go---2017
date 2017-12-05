package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const inputFile = "input_day04.txt"

//ReadInput reads the given puzzle input, splits at tabs and returns an array of integer array
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

//SortString sorts the letters in the given string w and returns the sorted one
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {

	input := ReadInput(inputFile)
	sum := len(input)

	for _, l := range input {
		hist := make(map[string]int)
		for _, w := range strings.Fields(l) {
			if _, ok := hist[w]; ok {
				sum--
				break

			} else {
				hist[w]++
			}
		}
	}
	fmt.Println("Part 1: ", sum)

	sum = len(input)
	for _, l := range input {
		hist := make(map[string]int)
		for _, w := range strings.Fields(l) {
			s := SortString(w)
			if _, ok := hist[s]; ok {
				sum--
				break
			} else {
				hist[s]++
			}
		}
	}
	fmt.Println("Part 2: ", sum)
}
