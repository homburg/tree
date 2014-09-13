package tree

import (
	"fmt"
	"github.com/fatih/color"
	"sort"
	"strings"
)

type node struct {
	nodes map[string]*node
}

const (
	TEE   = "├── "
	PIPE  = "│   "
	BEND  = "└── "
	WIDTH = len(PIPE)
)

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

func (n *node) Format(indent string, t *tree) string {
	var lines []string
	if n.isLeaf() || len(n.nodes) == 0 {
		return ""
	} else {
		keys := make([]string, len(n.nodes))
		i := 0
		for key := range n.nodes {
			keys[i] = key
			i++
		}

		sort.Strings(keys)

		i = len(n.nodes) - 1
		for _, key := range keys {
			c := n.nodes[key]
			newLine := indent[:len(indent)-WIDTH]
			if i == 0 {
				newLine += BEND
				indent = indent[:len(indent)-WIDTH] + "    "
			} else {
				newLine += TEE
			}
			newLine += fmt.Sprintf(t.NodeFormat, key)
			lines = append(
				lines,
				newLine,
			)

			if !c.isLeaf() {
				lines = append(
					lines,
					c.Format(indent+PIPE, t),
				)
			}
			i -= 1
		}
		return strings.Join(lines, "\n")
	}
	return ""
}

type tree struct {
	separator  string
	root       *node
	NodeFormat string
}

func (g *tree) Format() string {
	return ".\n" + g.root.Format("│   ", g) + "\n"
}

func New(separator string) *tree {
	return &tree{
		separator,
		&node{},
		color.BlueString("%%s"),
	}
}

func (g *tree) Eat(line string) {
	g.root.eat(strings.Split(line, g.separator))
}

func (g *tree) EatLines(lines []string) {
	for _, line := range lines {
		if line != "" {
			g.Eat(line)
		}
	}
}
