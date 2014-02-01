package main

import (
	"fmt"
	"github.com/roovo/tgl/gitlab"
)

func main() {
	gitlab := gitlab.NewGitlab("http://demo.gitlab.com")

	contents, err := gitlab.Projects()

	fmt.Printf("%n: %s\n", err, string(contents))
}
