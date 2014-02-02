package gitlab

import (
	"encoding/json"
	"fmt"
	"github.com/howeyc/gopass"
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

func NewGitlab() *Gitlab {
	g := &Gitlab{}

	err := g.load()
	if err != nil {
		var username, password string

		fmt.Printf("gitlab url: ")
		fmt.Scanln(&g.Url)
		fmt.Printf("username: ")
		fmt.Scanln(&username)
		fmt.Printf("password: ")
		password = string(gopass.GetPasswd())

		g.Login(username, password)
	}
	return g
}

func (g *Gitlab) Login(login string, password string) (err error) {
	var session *Session

	params := &map[string]string{
		"login":    login,
		"password": password,
	}

	err = g.postTo("/session", params, &session)
	if err != nil {
		return
	}

	g.Token = session.PrivateToken
	g.save()

	return
}

func (g *Gitlab) Projects() (projects []*Project, err error) {
	err = g.getFrom("/projects", &projects)

	return
}

func (g *Gitlab) apiUrlFor(path string) string {
	return g.Url + "/api/v3" + path + "?private_token=" + g.Token
}

func (g *Gitlab) getFrom(path string, out interface{}) (err error) {
	url := g.apiUrlFor(path)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return json.Unmarshal(contents, &out)

	return
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

func (g *Gitlab) postTo(path string, params *map[string]string, out interface{}) (err error) {
	request_url := g.apiUrlFor(path)

	values := make(url.Values)

	for k, v := range *params {
		values.Set(k, v)
	}

	resp, err := http.PostForm(request_url, values)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return json.Unmarshal(contents, &out)
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
