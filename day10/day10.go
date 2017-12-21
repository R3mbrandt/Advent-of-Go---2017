package main

import (
	"fmt"
)

func reverse(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func selElements(pos, number int, list []int) []int {
	out := make([]int, number)

	for i := 0; i < number; i++ {
		out[i] = list[(pos+i)%len(list)]
	}
	return out
}

func putElements(pos int, elements []int, list []int) {
	for i, e := range elements {
		list[(pos+i)%len(list)] = e
	}
}

func generateList(length int) []int {
	out := make([]int, length)
	for i := 0; i < length; i++ {
		out[i] = i
	}
	return out
}

func main() {
	var currPos = 0
	var skipSize = 0
	var length = []int{199, 0, 255, 136, 174, 254, 227, 16, 51, 85, 1, 2, 22, 17, 7, 192} // puzzle Input Part 1
	var byteStream = "199,0,255,136,174,254,227,16,51,85,1,2,22,17,7,192"                 // puzzle input Part 2

	list := generateList(256)

	for _, l := range length {
		if l > len(list) {
			continue
		}
		tmp := selElements(currPos, l, list)
		reverse(tmp)
		putElements(currPos, tmp, list)
		currPos += l + skipSize
		skipSize++
	}
	fmt.Println("Part 1:", list[0]*list[1])

	// PART 2
	currPos = 0
	skipSize = 0
	list = generateList(256) // make a new list
	//prepare Input for Part 2
	length = nil
	for _, r := range byteStream {
		length = append(length, int(r))
	}
	length = append(length, []int{17, 31, 73, 47, 23}...) // append suffix

	for i := 0; i < 64; i++ {
		for _, l := range length {
			if l > len(list) {
				continue
			}
			tmp := selElements(currPos, l, list)
			reverse(tmp)
			putElements(currPos, tmp, list)
			currPos += l + skipSize
			skipSize++
		}
	}
	denseHash := make([]int, 16)

	for i := 0; i < 16; i++ {
		tmp := list[i*16 : i*16+16]
		sum := tmp[0]
		for _, e := range tmp[1:] {
			sum ^= e
		}
		denseHash[i] = sum
	}
	fmt.Printf("Part 2: ")
	for _, e := range denseHash {
		fmt.Printf("%x", e)
	}
	fmt.Println()
}
