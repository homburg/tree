# Tree

[![Build Status](https://travis-ci.org/homburg/tree.svg?branch=master)](https://travis-ci.org/homburg/tree)

Create an ascii tree (like `$ tree`) from any text
using a custom delimiter.

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
		"first",
		"first-second.txt",
		"third",
	})

	fmt.Println(t.Format())
	// Output:
	// .
	// ├── first
	// │   └── second.txt
	// └── third.txt
}
```

## TODO

- test
- cli
  - CaaP (cli as a pipe)

## LICENSE

MIT 2014 Thomas B Homburg

