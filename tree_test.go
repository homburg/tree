package tree

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func ExampleTree() {
	file, err := ioutil.ReadFile("files.txt")
	if err != nil {
		log.Fatal(err)
	}

	g := New("/")
	lines := strings.Split(string(file), "\n")

	g.EatLines(lines)

	fmt.Print(g.Format())
	// Output:
	// .
	// ├── 1
	// │   ├── 2
	// │   │   └── 3
	// │   │       ├── 4
	// │   │       │   ├── 5
	// │   │       │   └── fisk2.txt
	// │   │       └── fisk2.txt
	// │   ├── 3
	// │   │   ├── 4
	// │   │   │   ├── 5
	// │   │   │   └── fisk2.txt
	// │   │   └── fisk2.txt
	// │   ├── 5
	// │   │   ├── 4
	// │   │   │   ├── 3
	// │   │   │   │   ├── 2
	// │   │   │   │   └── fisk.txt
	// │   │   │   └── fisk.txt
	// │   │   ├── fisk.txt
	// │   │   └── fisk2.txt
	// │   └── fisk.txt
	// └── fisk.txt
	//
}

func TestShallowTree(t *testing.T) {
	lines := []string{
		"one",
		"other",
		"this",
	}

	expected := `.
├── one
├── other
└── this
`
	tr := New("/")
	tr.EatLines(lines)

	output := tr.Format()

	errorFormat := `Expected
===
%s===

Got
===
%s===`

	if output != expected {
		t.Error("fisk...")
		t.Errorf(
			errorFormat,
			expected,
			output,
		)
	}
}
