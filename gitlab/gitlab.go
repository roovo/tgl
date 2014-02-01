package gitlab

import (
	"io/ioutil"
	"net/http"
)

type Gitlab struct {
	url string ""
}

func NewGitlab(a_url string) *Gitlab {
	return &Gitlab{
		url: a_url,
	}
}

func (g *Gitlab) Url() string {
	return g.url + "/api/v3"
}

func (g *Gitlab) Projects() ([]byte, error) {
	url := g.Url() + "/projects"

	resp, err := http.Get(url)
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
