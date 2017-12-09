package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	label    string
	weight   int
	parent   *Node
	children []*Node
}

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

func newNode(label string, weight int) *Node {
	n := new(Node)
	n.label = label
	n.weight = weight
	n.children = []*Node{}
	n.parent = nil

	return n
}

//just for a pretty Print while in the developement phase
func (n *Node) String() string {
	if n == nil {
		return "()"
	}
	s := ""
	if n.parent != nil {
		s += fmt.Sprint(n.parent.label) + "->"
	}
	s += fmt.Sprint(n.label)
	s += "(" + fmt.Sprint(n.weight) + ")"
	if len(n.children) > 0 {
		s += " | "
		for _, c := range n.children {
			s += c.String() + " "
		}
		s += "|"
	}
	return s
}

func (n *Node) addChild(c *Node) {
	n.children = append(n.children, c)
	c.parent = n
}

func (n *Node) childWeights() map[string]int {
	weights := make(map[string]int)
	if len(n.children) > 0 {
		for _, child := range n.children {
			weights[child.label] = child.totalWeight()
		}
	}
	return weights
}

func (n *Node) totalWeight() int {
	w := n.weight
	for _, c := range n.children {
		w += c.totalWeight()
	}
	return w
}

func (n *Node) balanced() bool {
	childweights := make(map[int]bool)
	for _, value := range n.childWeights() {
		childweights[value] = true
	}
	return len(childweights) == 1
}

func main() {
	input := ReadInput("input_day07.txt")
	re := regexp.MustCompile(`(.*)\s?\((\d+)\)\s?(?:->)?\s?(.*)`)
	nodes := make(map[string]*Node)

	for _, line := range input {
		result := re.FindStringSubmatch(line)
		label := strings.TrimSpace(result[1])
		weight, _ := strconv.Atoi(result[2])
		n := newNode(label, weight)
		nodes[label] = n
	}

	for _, line := range input {
		result := re.FindStringSubmatch(line)
		if result[3] != "" {
			parentLabel := strings.TrimSpace(result[1])
			parent := nodes[parentLabel]
			for _, childLabel := range strings.Split(result[3], ", ") {
				child := nodes[childLabel]
				parent.addChild(child)
			}
		}
	}
	rootLabel := ""
	for key, value := range nodes {
		if value.parent == nil {
			rootLabel = key
			break
		}
	}
	fmt.Println("Part 1:", rootLabel)

	root := nodes[rootLabel]

	for {
		childweights := root.childWeights()
		weights := make(map[int]int)
		var imbalancedWeight int
		var imbalancedLabel string

		for _, value := range childweights {
			weights[value]++
		}
		for key, value := range weights {
			if value == 1 {
				imbalancedWeight = key
				break
			}
		}
		for key, value := range childweights {
			if value == imbalancedWeight {
				imbalancedLabel = key
				break
			}
		}
		imbalancedNode := nodes[imbalancedLabel]
		if imbalancedNode.balanced() {
			fmt.Println(root.label, root.weight, childweights, imbalancedNode.label, imbalancedNode.weight)
			diffs := []int{}
			for key := range weights {
				diffs = append(diffs, key)
			}
			fmt.Println("Part 2:", imbalancedNode.weight-int((math.Abs(float64(diffs[0]-diffs[1])))))
			break
		} else {
			root = imbalancedNode
		}
	}
}
