package tree

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io"
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

func (n *node) eat(segments []string) error {
	if len(segments) == 0 {
		return nil
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

	return nil
}

func (n *node) Format(indent string, t *tree) string {
	var lines []string
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

		i = len(n.nodes) - 1
		for _, key := range keys {
			c := n.nodes[key]
			if !(!t.KeepLeaves && c.isLeaf()) {
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
					childLine := c.Format(indent+PIPE, t)
					if childLine != "" {
						lines = append(
							lines,
							childLine,
						)
					}
				}
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
	KeepLeaves bool
}

func (g *tree) Format() string {
	keys := make([]string, len(g.root.nodes))

	i := 0
	for key := range g.root.nodes {
		keys[i] = key
		i++
	}

	sort.Strings(keys)

	var treeStr string
	for _, key := range keys {
		n := g.root.nodes[key]
		treeStr += fmt.Sprintf(g.NodeFormat, key) + "\n"

		if !n.isLeaf() {
			treeStr += n.Format(PIPE, g) + "\n"
		}
	}

	return treeStr
}

func New(separator string) *tree {
	return &tree{
		separator,
		&node{},
		color.BlueString("%%s"),
		true,
	}
}

func (g *tree) eat(line string) {
	g.root.eat(strings.Split(line, g.separator))
}

func (g *tree) ReadAll(rdr io.Reader) error {
	scanner := bufio.NewScanner(rdr)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			g.eat(line)
		}
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		return err
	}

	return nil
}
