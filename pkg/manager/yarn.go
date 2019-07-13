package manager

import (
	"fmt"
	"net/http"
	"os/exec"
)

const YARN = "yarn"

type Yarn struct{}

func (m *Yarn) Name() string {
	return YARN
}

func (m *Yarn) Exists(name string) bool {
	resp, _ := m.callApi(fmt.Sprintf("/%s.json", name))
	return resp.StatusCode == 200
}

func (*Yarn) Add(name string) {
	exec.Command("yarn", "global", "add", "--silent", name)
}

func (*Yarn) callApi(path string) (resp *http.Response, err error) {
	endPoint := "https://replicate.npmjs.com"
	return http.Get(endPoint + path)
}