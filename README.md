# Tree

[![Build Status](https://travis-ci.org/homburg/tree.svg?branch=master)](https://travis-ci.org/homburg/tree)

Create an ascii tree (like `$ tree`) from any text
using a custom delimiter.

## Cli

https://github.com/homburg/tree/tree/master/cmd/tree

```bash
$ go get github.com/homburg/tree/cmd/tree
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/homburg/tree"
)

func main() {

	t := tree.New("-")

	t.EatLines([]string{
		"root"
		"root-first",
		"root-first-second.txt",
		"third",
	})

	fmt.Println(t.Format())
	// Output:
	// root
	// └─── first
	//      └── second.txt
	// third.txt
}
```

## TODO

- [x] test
- [x] cli
  - [x] CaaP (cli as a pipe)

## LICENSE(S)

MIT 2014 Thomas B Homburg

[LICENSE.md](LICENSE.md)
[LICENSES.md](LICENSES.md)

