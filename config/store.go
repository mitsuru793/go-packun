package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/mitsuru793/go-packun/pkg"
	"github.com/mitsuru793/go-packun/util"
)

type Store struct {
	Path string
}

func (s *Store) ReadFile() *os.File {
	return util.EnsureFile(s.Path)
}

func (s *Store) Read() (*pkg.Store, error) {
	p := s.expandPath()
	util.EnsureFile(p)

	var store pkg.Store
	_, err := toml.DecodeFile(p, &store)
	return &store, err
}

func (s *Store) expandPath() string {
	p := filepath.Clean(s.Path)
	return util.ExpandHome(p)
}

func (s *Store) exist() bool {
	return util.Exists(s.Path)
}
