package main

import (
	"fmt"
	"github.com/roovo/tgl/gitlab"
)

func main() {
	gitlab := gitlab.NewGitlab("http://demo.gitlab.com")

	gitlab.Login("test@test.com", "123456")

	projects, _ := gitlab.Projects()

	for _, project := range projects {
		fmt.Printf("%s\n", project.Name)
	}
}
