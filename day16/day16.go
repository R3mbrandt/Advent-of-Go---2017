package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		out = append(out, strings.Split(scanner.Text(), ",")...)
	}
	return out
}

func findPos(elem string, progs []string) int {
	for n, p := range progs {
		if p == elem {
			return n
		}
	}
	return 0 // should never reach this - in production return (int, error) tuple
}

func dance(instr string, progs []string) {
	switch instr[0] {
	case 's':
		if spins, err := strconv.Atoi(instr[1:]); err == nil {
			tmp := append(progs[len(progs)-spins:], progs[:len(progs)-spins]...)
			for n := range progs {
				progs[n] = tmp[n]
			}
		}
	case 'x':
		tmp := strings.Split(instr[1:], "/")
		a, _ := strconv.Atoi(tmp[0])
		b, _ := strconv.Atoi(tmp[1])
		progs[a], progs[b] = progs[b], progs[a]
	case 'p':
		tmp := strings.Split(instr[1:], "/")
		a := findPos(tmp[0], progs)
		b := findPos(tmp[1], progs)
		progs[a], progs[b] = progs[b], progs[a]
	}
}

func main() {
	var progs []string
	input := ReadInput("input_day16.txt")

	for i := 0; i < 16; i++ {
		progs = append(progs, string(i+97))
	}
	for _, instr := range input {
		dance(instr, progs)
	}

	fmt.Println("Part 1:", strings.Join(progs, ""))

	// Brute Force Attempt will NOT WORK :-(
	/*for i := 0; i < 1000000000; i++ {
		for _, instr := range input {
			dance(instr, progs)
		}
		fmt.Printf("%8d\n", i)
	}*/

	// Have to try another way...

	seen := make(map[string]int)
	var pos = 1
	seen[strings.Join(progs, "")] = 1

	for {
		for _, instr := range input {
			dance(instr, progs)
		}
		join := strings.Join(progs, "")
		pos++
		if _, ok := seen[join]; ok {
			break
		}
		seen[join] = pos
	}

	lastPos := 1000000000 % len(seen)
	for key, value := range seen {
		if value == lastPos {
			fmt.Println("Part 2:", key)
		}
	}

}
