package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/homburg/tree"
	"github.com/homburg/tree/cmd/tree/Godeps/_workspace/src/github.com/andrew-d/go-termutil"
	"github.com/homburg/tree/cmd/tree/Godeps/_workspace/src/github.com/kardianos/osext"
)

func removePathFromEnv(removePath string) {
	envPath := os.Getenv("PATH")
	paths := strings.Split(envPath, string(os.PathListSeparator))
	var newPaths []string
	for _, path := range paths {
		if path == removePath || (path+"/") == removePath {
			// Do nothing
		} else {
			newPaths = append(newPaths, path)
		}
	}

	os.Setenv("PATH", strings.Join(newPaths, string(os.PathListSeparator)))
}

var trimLeaves []bool

func anyTrue(val0 bool, vals ...bool) bool {
	if val0 {
		return true
	}
	for _, v := range vals {
		if v {
			return true
		}
	}

	return false
}

func init() {

	trimLeaves = []bool{false, false, false}

	flag.BoolVar(&trimLeaves[0], "trim-leaves", false, "")
	flag.BoolVar(&trimLeaves[1], "-trim-leaves", false, "")
	flag.BoolVar(&trimLeaves[2], "trim", false, "")
}

func main() {
	if termutil.Isatty(os.Stdin.Fd()) {

		// Remove the current path from paths
		// and try to resolve the original tree command
		path, err := osext.ExecutableFolder()
		removePathFromEnv(path)

		whichTree, err := exec.LookPath("tree")
		if err == nil {
			syscall.Exec(whichTree, os.Args, os.Environ())
		}
		return
	}

	// pipe it!
	flag.Parse()

	separator := flag.Arg(0)
	if "" == separator {
		separator = string(os.PathSeparator)
	}

	log.Printf("Using separator: \"%s\"\n", separator)

	t := tree.New(separator)
	if !termutil.Isatty(os.Stdout.Fd()) {
		t.NodeFormat = "%s"
	}
	t.KeepLeaves = !(anyTrue(false, trimLeaves...))

	err := t.ReadAll(os.Stdin)

	if nil != err {
		log.Fatal(err)
	}

	os.Stdout.WriteString(t.Format() + "\n")
}
