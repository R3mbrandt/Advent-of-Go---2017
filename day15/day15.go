package main

import (
	"fmt"
)

const (
	startA = 277
	startB = 349

	factorA = 16807
	factorB = 48271

	divisor = 2147483647

	rounds = 40000000
)

func main() {
	var matches int
	var genA uint64 = startA
	var genB uint64 = startB

	for i := 0; i < rounds; i++ {
		genA = genA * factorA % divisor
		genB = genB * factorB % divisor
		if (genA & 0xFFFF) == (genB & 0xFFFF) { // 0xFFFF == 0b00000000000000001111111111111111 --- wir wollen nur die letzten 16bit vergleichen
			matches++ // und maskieren den Rest mit binären AND aus
		}
	}
	fmt.Println("Part 1:", matches)

	matches = 0
	genA = startA
	genB = startB
	for i := 0; i < rounds/8; i++ {
		for {
			genA = genA * factorA % divisor
			if genA%4 == 0 {
				break
			}
		}
		for {
			genB = genB * factorB % divisor
			if genB%8 == 0 {
				break
			}
		}
		if (genA & 0xFFFF) == (genB & 0xFFFF) {
			matches++
		}
	}
	fmt.Println("Part 2:", matches)
}
