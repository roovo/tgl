package main

import (
	"fmt"
	"github.com/roovo/tgl/gitlab"
)

func main() {
	gitlab := gitlab.NewGitlab("http://demo.gitlab.com")

	gitlab.Login("test@test.com", "123456")

	fmt.Printf("token: %s\n", gitlab.Token)
}
