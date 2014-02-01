package gitlab

import (
	"io/ioutil"
	"net/http"
)

func Projects() ([]byte, error) {
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
