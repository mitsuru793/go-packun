package util

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

func EnsureFile(path string) *os.File {
	p := ExpandHome(path)
	if Exists(p) {
		f, _ := os.OpenFile(p, os.O_WRONLY, 0644)
		return f
	}

	dir := filepath.Dir(p)
	if !Exists(dir) {
		os.Mkdir(dir, os.ModePerm)
	}
	f, _ := os.OpenFile(p, os.O_CREATE, 0644)
	return f
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func ExpandHome(path string) string {
	usr, _ := user.Current()
	return strings.Replace(path, "~", usr.HomeDir, 1)
}

func SaveToml(f *os.File, data interface{}) error {
	if err := toml.NewEncoder(f).Encode(data); err != nil {
		return err
	}
	return nil
}
