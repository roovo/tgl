package main

import (
	"fmt"
	"github.com/roovo/tgl/gitlab"
)

func main() {
	gitlab := gitlab.NewGitlab("http://demo.gitlab.com")

	contents, err := gitlab.Login("test@test.com", "123456")

	fmt.Printf("%n: %s\n", err, string(contents))
}
