package tree

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

func (n *node) Format(indent string) string {
	var lines []string
	if n.isLeaf() || len(n.nodes) == 0 {
		panic("")
	} else {
		// for key, c := range n.nodes {
		i := len(n.nodes) - 1
		for key, c := range n.nodes {
			newLine := indent[:len(indent)-8]
			if i == 0 {
				newLine += "└── "
				indent = indent[:len(indent)-8] + "    "
			} else if len(indent) == 8 && i == (len(n.nodes)-1) {
				newLine += "."
			} else {
				newLine += "├── "
			}
			newLine += key
			lines = append(
				lines,
				newLine,
			)

			if !c.isLeaf() {
				lines = append(
					lines,
					c.Format(indent+"│   "),
				)
			}
			i -= 1
		}
		return strings.Join(lines, "\n")
	}
}

type tree struct {
	separator string
	root      *node
}

func (g *tree) String() string {
	return g.root.String()
}

func (g *tree) Format() string {
	return g.root.Format("│   ")
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
		"1.4",
		"1.4.3",
		"1.4.4",
		"2",
		"17",
	}

	g := newTree(".")

	for _, line := range lines {
		g.eat(line)
	}

	fmt.Println(g.Format())
}
