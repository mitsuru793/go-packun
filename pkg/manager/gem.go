package manager

import (
	"fmt"
	"net/http"
	"github.com/mitsuru793/go-packun/util"
)

const GEM = "gem"

type Gem struct{}

func (m *Gem) Name() string {
	return GEM
}

func (m *Gem) Exists(name string) bool {
	resp, _ := m.callApi(fmt.Sprintf("/gems/%s.json", name))
	return resp.StatusCode == 200
}

func (*Gem) Add(name string) {
	util.RunCommand("gem install " + name)
}

func (*Gem) Remove(name string) {
	util.RunCommand("gem uninstall " + name)
}

func (*Gem) callApi(path string) (resp *http.Response, err error) {
	endPoint := "https://rubygems.org/api/v1"
	return http.Get(endPoint + path)
}
