package gitlab

import (
	"io/ioutil"
	"net/http"
	"net/url"
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

func (g *Gitlab) Login(login string, password string) ([]byte, error) {
	request_url := g.Url() + "/session"

	values := make(url.Values)
	values.Set("login", login)
	values.Set("password", password)

	resp, err := http.PostForm(request_url, values)
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
