package main

import (
	"github.com/homburg/tree/cmd/tree/Godeps/_workspace/src/bitbucket.org/kardianos/osext"
	"github.com/homburg/tree/cmd/tree/Godeps/_workspace/src/github.com/andrew-d/go-termutil"
	"github.com/homburg/tree/cmd/tree/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/homburg/tree"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
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

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "trim-leaves, trim",
			Usage: "Trim leaves.",
		},
	}
	app.Action = func(c *cli.Context) {
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
		input, err := ioutil.ReadAll(os.Stdin)
		if nil != err {
			log.Fatal(err)
		}

		lines := strings.Split(string(input), "\n")

		separator := c.Args().First()
		if "" == separator {
			separator = string(os.PathSeparator)
		}

		log.Printf("Using separator: \"%s\"\n", separator)

		t := tree.New(separator)
		if !termutil.Isatty(os.Stdout.Fd()) {
			t.NodeFormat = "%s"
		}
		t.KeepLeaves = !c.Bool("trim-leaves")
		t.EatLines(lines)
		os.Stdout.WriteString(t.Format() + "\n")
	}
	app.Run(os.Args)
}
