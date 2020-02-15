package manager

import (
	"fmt"
	"net/http"
	"github.com/mitsuru793/go-packun/util"
)

const COMPOSER = "composer"

type Composer struct{}

func (m *Composer) Name() string {
	return COMPOSER
}

func (m *Composer) Exists(name string) bool {
	resp, _ := m.callApi(fmt.Sprintf("/p/%s.json", name))
	return resp.StatusCode == 200
}

func (*Composer) Add(name string) {
	util.RunCommand("composer global require " + name)
}

func (*Composer) Remove(name string) {
	util.RunCommand("composer global remove " + name)
}

func (*Composer) callApi(path string) (resp *http.Response, err error) {
	endPoint := "https://repo.packagist.org"
	return http.Get(endPoint + path)
}
