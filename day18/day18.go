package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//ReadInput reads line by line from filename and returns a map of ints
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
		out = append(out, scanner.Text())
	}
	return out
}

func ParseIntstruction(instr string, registers map[string]int) (string, string, int) {
	tmp := strings.Split(instr, " ")
	if len(tmp) < 3 {
		tmp = append(tmp, "0")
	}
	rv, e := strconv.Atoi(tmp[2])
	if e != nil {
		rv = registers[tmp[2]]
	}

	return tmp[0], tmp[1], rv
}

func ExecuteInstruction(instr, lv string, rv int, registers map[string]int, instrPointer *int) (error, bool) {
	switch instr {
	case "snd":
		registers["snd"] = registers[lv]
	case "set":
		registers[lv] = rv
	case "add":
		registers[lv] += rv
	case "mul":
		registers[lv] *= rv
	case "mod":
		registers[lv] %= rv
	case "rcv":
		if registers[lv] != 0 {
			fmt.Println("Frequency of last played Sound", registers["snd"])
			return nil, true
		}
	case "jgz":
		if registers[lv] > 0 {
			*instrPointer += rv - 1
		}
	default:
		return errors.New("Unknown Instruction"), false
	}
	return nil, false
}

func main() {
	var instrPointer int
	registers := make(map[string]int)
	input := ReadInput("input_day18.txt")
	for {
		instr := input[instrPointer]
		instrPointer++
		ins, lv, rv := ParseIntstruction(instr, registers)
		e, end := ExecuteInstruction(ins, lv, rv, registers, &instrPointer)
		if e == nil && end == true {
			fmt.Println(registers)
			break
		}
	}
}
