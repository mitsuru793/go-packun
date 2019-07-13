package manager

import (
	"fmt"
	"net/http"
	"os/exec"
)

const BREW string = "brew"

type Brew struct{}

func (m *Brew) Name() string {
	return BREW
}

func (m *Brew) Exists(name string) bool {
	resp, _ := m.callApi(fmt.Sprintf("/formula/%s.json", name))
	return resp.StatusCode == 200
}

func (*Brew) Add(name string) {
	exec.Command("brew", "install", name)
}

func (*Brew) callApi(path string) (resp *http.Response, err error) {
	endPoint := "https://formulae.brew.sh/api"
	return http.Get(endPoint + path)
}
