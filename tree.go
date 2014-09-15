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

func (n *node) Format(out chan<- string, indent string, t *tree) {
	if n.isLeaf() || len(n.nodes) == 0 {
		// pass
	} else {
		keys := make([]string, len(n.nodes))
		i := 0
		for key := range n.nodes {
			keys[i] = key
			i++
		}

		sort.Strings(keys)

		var outs [](chan string)
		i = len(n.nodes) - 1
		for _, key := range keys {
			newLineOut := make(chan string)
			outs = append(outs, newLineOut)

			go func(out chan<- string, key string, indent string, i int, child *node) {
				newLine := indent[:len(indent)-WIDTH]
				if i == 0 {
					newLine += BEND
					indent = indent[:len(indent)-WIDTH] + "    "
				} else {
					newLine += TEE
				}
				newLine += fmt.Sprintf(t.NodeFormat, key)

				if !child.isLeaf() {
					subOut := make(chan string)
					go child.Format(subOut, indent+PIPE, t)
					newLine += "\n" + <-subOut
				}

				out <- newLine
				close(out)
			}(newLineOut, key, indent, i, n.nodes[key])
			i--

		}

		lines := make([]string, len(outs))
		i = 0
		var output string
		for _, c := range outs {
			output = <-c
			lines[i] = output
			i++
		}

		out <- strings.Join(lines, "\n")
		close(out)
	}
}

type tree struct {
	separator  string
	root       *node
	NodeFormat string
}

func (g *tree) Format() string {
	out := make(chan string)
	go g.root.Format(out, "│   ", g)
	return ".\n" + (<-out) + "\n"
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
