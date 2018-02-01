package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const inputstring = `snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d`

func execInstr(instr string, registers map[string]int, snd, rcv chan int) bool {
	fields := strings.Fields(instr)
	var out int
	switch fields[0] {
	case "snd":
		if v, e := strconv.Atoi(fields[1]); e == nil {
			out = v
		} else {
			out = registers[fields[1]]
		}
		select {
		case snd <- out:
			return true
		case <-time.After(500 * time.Millisecond):
			return false
		}
	case "rcv":
		select {
		case val := <-rcv:
			registers[fields[1]] = val
		case <-time.After(500 * time.Millisecond):
			return false
		}
		return false
	}
	return false
}

func main() {
	var counter int
	regs1 := make(map[string]int)
	regs2 := make(map[string]int)
	input := strings.Split(inputstring, "\n")
	var snd chan int
	var rcv chan int
	snd = make(chan int, 2000)
	rcv = make(chan int, 2000)

	go func() {
		regs1["p"] = 1
		for i := 0; i < len(input); i++ {
			if execInstr(input[i], regs1, snd, rcv) {
				counter++
			}
		}
	}()
	func() {
		regs2["p"] = 0
		for i := 0; i < len(input); i++ {
			execInstr(input[i], regs2, rcv, snd)
		}
	}()
	fmt.Println(counter)
}
