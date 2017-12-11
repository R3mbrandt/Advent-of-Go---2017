package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

//ReadInput reads line by line from filename and returns an string array
func ReadInput(filename string) []string {
	var out []string
	fn, err := os.Open(filename) // open "filename" and panics at any error
	if err != nil {
		panic("File Error: " + err.Error())
	}
	defer fn.Close()

	scanner := bufio.NewScanner(fn)
	for scanner.Scan() {
		out = append(out, scanner.Text()) // reads one line of the file as string and appends to the string array

	}
	return out
}

func checkCondition(registers map[string]int, register, cond string, value int) bool {
	switch cond {
	case "<":
		return registers[register] < value
	case ">":
		return registers[register] > value
	case "<=":
		return registers[register] <= value
	case ">=":
		return registers[register] >= value
	case "==":
		return registers[register] == value
	case "!=":
		return registers[register] != value
	}
	return false
}

func doInstruction(registers map[string]int, register, instr string, value int) {
	if _, ok := registers[register]; !ok {
		registers[register] = 0
	}
	switch instr {
	case "inc":
		registers[register] += value
	case "dec":
		registers[register] -= value
	}
}

func max(array []int) int {
	var max = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func findBiggestValue(registers map[string]int) int {
	values := make([]int, len(registers))
	for _, value := range registers {
		values = append(values, value)
	}
	return max(values)
}

func main() {
	input := ReadInput("input_day08.txt")
	re := regexp.MustCompile(`(.+) (inc|dec) (-?\d+) if (.+) ([<>=!]+) (-?\d+)`)
	registers := make(map[string]int)
	highestValue := 0

	for _, line := range input {
		result := re.FindStringSubmatch(line)
		value, _ := strconv.Atoi(result[6])
		if checkCondition(registers, result[4], result[5], value) {
			value, _ := strconv.Atoi(result[3])
			doInstruction(registers, result[1], result[2], value)
			if findBiggestValue(registers) > highestValue {
				highestValue = findBiggestValue(registers)
			}
		}
	}
	fmt.Println("Part 1:", findBiggestValue(registers))
	fmt.Println("Part 2:", highestValue)
}
