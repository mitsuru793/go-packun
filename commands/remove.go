package commands

import (
	"fmt"
	"github.com/mitsuru793/go-packun/config"
	"github.com/mitsuru793/go-packun/pkg"
	"github.com/urfave/cli"
)

func NewRemove() cli.Command {
	return cli.Command{
		Name:    "remove",
		Aliases: []string{"r"},
		Usage:   "remove a package",
		Action:  removeAction,
	}
}

func removeAction(c *cli.Context) error {
	return action(c, func(c *cli.Context, config *config.Config) error {
		pType, err := askPackageManager()
		if err != nil {
			return err
		}
		pkgManager := pkg.NewPackageManager(pType)

		removePkg, err := askPackageName(pkgManager)
		if err != nil {
			return err
		}

		store, err := config.Store.Read()
		if err != nil {
			return err
		}

		manager := pkg.NewPackageManager(removePkg.Manager)
		manager.Remove(removePkg.Name)

		store.Packages.Remove(removePkg)
		store.Save(config.Store.ReadFile())

		fmt.Println("Removed:", removePkg.Name)

		return nil
	})
}
