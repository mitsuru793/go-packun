package manager

import (
	"fmt"
	"net/http"
	"github.com/mitsuru793/go-packun/util"
)

const BREW_CASK = "brew-cask"

type BrewCask struct{}

func (m *BrewCask) Name() string {
	return BREW_CASK
}

func (m *BrewCask) Exists(name string) bool {
	resp, _ := m.callApi(fmt.Sprintf("/cask/%s.json", name))
	return resp.StatusCode == 200
}

func (*BrewCask) Add(name string) {
	util.RunCommand("brew cask install " + name)
}

func (*BrewCask) callApi(path string) (resp *http.Response, err error) {
	endPoint := "https://formulae.brew.sh/api"
	return http.Get(endPoint + path)
}
