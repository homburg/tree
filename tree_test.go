package tree

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/renstrom/dedent"
)

func ExampleTree() {
	file, err := os.Open("files.txt")
	if err != nil {
		log.Fatal(err)
	}

	g := New("/")

	g.ReadAll(file)

	fmt.Print(g.Format())
	// Output:
	//
	// [34m1[0m
	// ├── [34m2[0m
	// │   └── [34m3[0m
	// │       ├── [34m4[0m
	// │       │   ├── [34m5[0m
	// │       │   └── [34mfisk2.txt[0m
	// │       └── [34mfisk2.txt[0m
	// ├── [34m3[0m
	// │   ├── [34m4[0m
	// │   │   ├── [34m5[0m
	// │   │   └── [34mfisk2.txt[0m
	// │   └── [34mfisk2.txt[0m
	// ├── [34m5[0m
	// │   ├── [34m4[0m
	// │   │   ├── [34m3[0m
	// │   │   │   ├── [34m2[0m
	// │   │   │   └── [34mfisk.txt[0m
	// │   │   └── [34mfisk.txt[0m
	// │   ├── [34mfisk.txt[0m
	// │   └── [34mfisk2.txt[0m
	// └── [34mfisk.txt[0m
	// [34mfisk.txt[0m
	//
}

func TestShallowTree(t *testing.T) {
	input := strings.NewReader(dedent.Dedent(`
		one
		other
		this
		`,
	))

	expected := dedent.Dedent(`[34mone[0m
		[34mother[0m
		[34mthis[0m
		`,
	)

	setup := func() *tree {
		return New("/")
	}
	RunFormatTestCase(input, expected, setup, t)
}

func RunFormatTestCase(input io.Reader, expected string, setup func() *tree, t *testing.T) {
	tr := setup()

	tr.ReadAll(input)

	output := tr.Format()

	errorFormat := dedent.Dedent(`Expected
		===
		%s===

		Got
		===
		%s===`,
	)

	if output != expected {
		t.Errorf(
			errorFormat,
			expected,
			output,
		)
	}
}

func TestNodeFormat(t *testing.T) {
	rdr := strings.NewReader(dedent.Dedent(`
		one
		other$retho
		this
		`,
	))

	expected := dedent.Dedent(`✓ one ⚡
		✓ other ⚡
		└── ✓ retho ⚡
		✓ this ⚡
		`,
	)

	setup := func() *tree {
		tr := New("$")
		tr.NodeFormat = "✓ %s ⚡"
		return tr
	}

	RunFormatTestCase(rdr, expected, setup, t)
}

func TestSingleNode(t *testing.T) {
	rdr := strings.NewReader("1")
	expected := "1\n"
	setup := func() *tree {
		tr := New(".")
		tr.NodeFormat = "%s"
		return tr
	}
	RunFormatTestCase(rdr, expected, setup, t)
}

func BenchmarkTreeFormat(b *testing.B) {
	file, err := os.Open("files.txt")
	if nil != err {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		t := New("/")
		t.ReadAll(file)
		t.Format()
	}
}

func BenchmarkHeavyTreeFormat(b *testing.B) {
	file, err := os.Open("data.csv")
	if nil != err {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		t := New(",")
		t.ReadAll(file)
		t.Format()
	}
}
