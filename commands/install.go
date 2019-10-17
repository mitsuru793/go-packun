package commands

import (
	"github.com/mitsuru793/go-packun/config"
	"github.com/urfave/cli"
	"github.com/mitsuru793/go-packun/pkg"
)

func NewInstall() cli.Command {
	return cli.Command{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "install packages from store",
		Action:  installAction,
	}
}

func installAction(c *cli.Context) error {
	return action(c, func(c *cli.Context, config *config.Config) error {
		store, err := config.Store.Read()
		if err != nil {
			return err
		}
		for _, p := range store.Packages {
			m := pkg.NewPackageManager(p.Manager)
			m.Add(p.Name)
		}
		return nil
	})
}
