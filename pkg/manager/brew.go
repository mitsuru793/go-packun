package manager

import (
	"fmt"
	"net/http"
	"github.com/mitsuru793/go-packun/util"
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
	util.RunCommand("brew install " + name)
}

func (*Brew) Remove(name string) {
	util.RunCommand("brew uninstall " + name)
}

func (*Brew) callApi(path string) (resp *http.Response, err error) {
	endPoint := "https://formulae.brew.sh/api"
	return http.Get(endPoint + path)
}
