package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func gitlabRequest() ([]byte, error) {
	resp, err := http.Get("http://demo.gitlab.com/api/v3/projects")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return contents, err
}

func main() {
	contents, err := gitlabRequest()

	fmt.Printf("%n: %s\n", err, string(contents))
}
