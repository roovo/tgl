package gitlab

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Gitlab struct {
	url   string
	token string
}

type Project struct {
	id                   int    `json:"id,omitempty`
	WebUrl               string `json:"web_url,omitempty`
	Name                 string `json:"name,omitempty`
	MergeRequestsEnabled bool   `json:"merge_requests_enabled,omitempty`
	lastActivity         string `json:"last_activity_at,omitempty`
}

type Session struct {
	PrivateToken string `json:"private_token,omitempty"`
}

func NewGitlab(a_url string) *Gitlab {
	return &Gitlab{
		url: a_url,
	}
}

func (g *Gitlab) apiUrlFor(path string) string {
	return g.url + "/api/v3" + path + "?private_token=" + g.token
}

func (g *Gitlab) Login(login string, password string) (err error) {
	request_url := g.apiUrlFor("/session")

	values := make(url.Values)
	values.Set("login", login)
	values.Set("password", password)

	resp, err := http.PostForm(request_url, values)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var session *Session

	err = json.Unmarshal(contents, &session)
	if err != nil {
		return
	}

	g.token = session.PrivateToken
	return
}

func (g *Gitlab) Projects() (projects []*Project, err error) {
	url := g.apiUrlFor("/projects")

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(contents, &projects)

	return
}
