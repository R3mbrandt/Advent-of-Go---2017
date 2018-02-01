package main

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

func splitToString(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}

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

func knotHash(input string) string {
	var currPos = 0
	var skipSize = 0
	list := generateList(256)
	var length []int

	for _, r := range input {
		length = append(length, int(r))
	}
	length = append(length, []int{17, 31, 73, 47, 23}...)
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
	s := ""
	for _, e := range denseHash {
		s += fmt.Sprintf("%02x", e)
	}
	return s
}

func markGroup(x, y int, grid *[128][128]int) {
	var queue []pos
	var neighbours = []pos{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}
	var tmpPos pos

	queue = append(queue, pos{x, y})

	for {
		p := queue[0]
		queue = queue[1:]

		grid[p.y][p.x] = 0

		for _, n := range neighbours {
			tmpPos.y = p.y + n.y
			tmpPos.x = p.x + n.x
			if tmpPos.y <= 127 && tmpPos.y >= 0 && tmpPos.x <= 127 && tmpPos.x >= 0 {
				if grid[tmpPos.y][tmpPos.x] == 1 {
					queue = append(queue, tmpPos)
				}
			}
		}

		if len(queue) == 0 {
			break
		}
	}
}

func main() {
	var input = "hxtvlmkl-"
	//var input = "flqrgnkx-"
	var used int
	var gridS [128]string
	var gridI [128][128]int
	var line string
	var groups int

	for counter := 0; counter < 128; counter++ {
		hash := knotHash(input + strconv.Itoa(counter))
		line = ""
		for _, r := range strings.Split(hash, "") {
			tmp, _ := strconv.ParseInt(r, 16, 8)
			used += bits.OnesCount(uint(tmp))
			line += fmt.Sprintf("%04b", tmp)
		}
		gridS[counter] = line
	}
	fmt.Println("Part 1:", used)

	// Part 2
	for y, l := range gridS {
		for x, e := range l {
			gridI[y][x] = int(e) - 48
		}
	}

	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			if gridI[y][x] == 1 {
				groups++
				markGroup(x, y, &gridI)
			}
		}
	}
	fmt.Println("Part 2:", groups)
}
