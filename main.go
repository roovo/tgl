package main

import (
  "github.com/roovo/tgl/gitlab"
	"fmt"
)

func main() {
	contents, err := gitlab.Projects()

	fmt.Printf("%n: %s\n", err, string(contents))
}
