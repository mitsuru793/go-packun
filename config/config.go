package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mitsuru793/go-packun/util"
)

type Config struct {
	Store *Store
}

func NewConfig() *Config {
	return &Config{
		&Store{
			Path: "~/.go-packun/store.toml",
		},
	}
}

func Load(path string) (*Config, error) {
	if !util.Exists(path) {
		return nil, fmt.Errorf("Not found file: %s", path)
	}
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) Save(f *os.File) error {
	return util.SaveToml(f, c)
}
