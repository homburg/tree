package main

import (
	"github.com/andrew-d/go-termutil"
	"github.com/codegangsta/cli"
	"github.com/homburg/tree"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

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
			whichTree, err := exec.LookPath("tree")
			if err != nil {
				syscall.Exec(whichTree, os.Args, os.Environ())
				return
			}
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
