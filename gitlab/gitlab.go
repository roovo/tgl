package gitlab

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

var configFilePath = filepath.Join(os.Getenv("HOME"), ".tglconfig")

type Gitlab struct {
	Url   string
	Token string
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
	g := &Gitlab{}
	g.load()
	return g
}

func (g *Gitlab) apiUrlFor(path string) string {
	return g.Url + "/api/v3" + path + "?private_token=" + g.Token
}

func (g *Gitlab) load() error {
	f, err := os.Open(configFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	for {
		if err := dec.Decode(g); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}

func (g *Gitlab) save() (err error) {
	f, err := os.Create(configFilePath)
	if err != nil {
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(g)
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

	g.Token = session.PrivateToken
	g.save()

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
