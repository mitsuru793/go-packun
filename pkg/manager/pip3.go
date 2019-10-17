package manager

import (
	"fmt"
	"net/http"
	"os/exec"
)

const PIP3 = "pip3"

type Pip3 struct{}

func (m *Pip3) Name() string {
	return PIP3
}

func (m *Pip3) Exists(name string) bool {
	resp, _ := m.callApi(fmt.Sprintf("/%s.json", name))
	return resp.StatusCode == 200
}

func (*Pip3) Add(name string) {
	exec.Command("pip3", "install", name).Run()
}

func (*Pip3) callApi(path string) (resp *http.Response, err error) {
	endPoint := "https://pypi.org/pypi"
	return http.Get(endPoint + path)
}
