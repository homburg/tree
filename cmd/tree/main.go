package main

import (
	"github.com/andrew-d/go-termutil"
	"github.com/homburg/tree"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
)

var trees = []string{
	"/usr/bin/tree",
	"/usr/local/bin/tree",
}

func main() {
	// Terminal?
	whichTree := ""
	didRun := false
	if termutil.Isatty(os.Stdin.Fd()) {
		for _, whichTree = range trees {
			_, err := os.Stat(whichTree)
			if os.IsNotExist(err) {
				continue
			}
			break
		}

		didRun = true
		syscall.Exec(whichTree, os.Args, os.Environ())
	}

	if !didRun {
		// pipe it!
		input, err := ioutil.ReadAll(os.Stdin)
		if nil != err {
			log.Fatal(err)
		}

		lines := strings.Split(string(input), "\n")

		separator := string(os.PathSeparator)
		if len(os.Args) > 1 {
			separator = os.Args[1]
		}

		t := tree.New(separator)
		t.EatLines(lines)
		os.Stdout.WriteString(t.Format() + "\n")
	}
}
