package pkg

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Store struct {
	Packages Packages
}

func (s *Store) Save(f *os.File) error {
	if err := toml.NewEncoder(f).Encode(s); err != nil {
		return err
	}
	return nil
}
