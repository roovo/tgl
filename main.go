package main

import "fmt"

func main() {
	var gitlab_url string
	var username string
	var password string

	fmt.Printf("gitlab url: ")
	fmt.Scanln(&gitlab_url)
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)
}
