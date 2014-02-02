package main

import (
	"fmt"
	"github.com/roovo/tgl/gitlab"
)

func main() {
	gitlab := gitlab.NewGitlab()

	projects, _ := gitlab.Projects()

	for _, project := range projects {
		fmt.Printf("%t: %s\n", project.MergeRequestsEnabled, project.Name)
	}
}
