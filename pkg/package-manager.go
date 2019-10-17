package pkg

import (
	"github.com/mitsuru793/go-packun/pkg/manager"
	"github.com/mitsuru793/go-packun/util"
)

var Types = []string{
	"brew", "pip", "gem", "composer", "yarn",
}

type PackageManager interface {
	Exists(name string) bool
	Add(name string)
	Name() string
}

func NewPackageManager(pType string) PackageManager {
	switch pType {
	case manager.BREW:
		return &manager.Brew{}
	case manager.BREW_CASK:
		return &manager.BrewCask{}
	case manager.YARN:
		return &manager.Yarn{}
	case manager.COMPOSER:
		return &manager.Composer{}
	case manager.PIP3:
		return &manager.Pip3{}
	case manager.GEM:
		return &manager.Gem{}
	default:
		panic("Invalid package type")
	}
}

func ExistsManager(pType string) bool {
	return util.Contains(Types, pType)
}
