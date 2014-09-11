package main

import (
	"fmt"
	"strings"
)

type node struct {
	nodes map[string]*node
}

func (n *node) isLeaf() bool {
	return nil == n.nodes || len(n.nodes) == 0
}

func (n *node) eat(segments []string) {
	if len(segments) == 0 {
		return
	}

	firstSegment := segments[0]
	var subNode *node
	if nil == n.nodes {
		subNode = &node{}
		n.nodes = make(map[string]*node)
		n.nodes[firstSegment] = subNode
	} else {
		if _, ok := n.nodes[firstSegment]; !ok {
			subNode = &node{}
			n.nodes[firstSegment] = subNode
		} else {
			subNode = n.nodes[firstSegment]
		}
	}

	if len(segments) > 1 {
		subNode.eat(segments[1:])
	}
}

func (n *node) String() string {
	if n.isLeaf() {
		return ""
	} else {
		return fmt.Sprint(n.nodes)
	}
}

type tree struct {
	separator string
	root      *node
}

func (g *tree) String() string {
	return g.root.String()
}

func newTree(separator string) *tree {
	return &tree{separator, &node{}}
}

func (g *tree) eat(line string) {
	g.root.eat(strings.Split(line, g.separator))
}

func main() {
	lines := []string{
		"1.2.fisk",
		"1.2.3.4",
		"1.2.3.fisk1",
		"1.2.3.fisk3",
		"1.2.3.fisk4",
		"1.2.3.fisk5",
		"1.2.3.fisk7",
		"1.2.3.fisk8",
	}

	g := newTree(".")

	for _, line := range lines {
		g.eat(line)
	}

	fmt.Println(g)
}
