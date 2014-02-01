package main

import (
	"fmt"
	"github.com/roovo/tgl/gitlab"
)

func main() {
	contents, err := gitlab.Projects()

	fmt.Printf("%n: %s\n", err, string(contents))
}
