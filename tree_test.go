package tree

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
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
}
