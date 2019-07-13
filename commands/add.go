package commands

import (
	"fmt"
	"strings"

	"github.com/mitsuru793/go-packun/config"
	"github.com/mitsuru793/go-packun/pkg"
	"github.com/mitsuru793/go-packun/util"
	"github.com/urfave/cli"
)

func NewAdd() cli.Command {
	return cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a package",
		Action:  addAction,
	}
}

func addAction(c *cli.Context) error {
	return action(c, func(c *cli.Context, config *config.Config) error {
		pType, err := askPackageManager()
		if err != nil {
			return err
		}
		pkgManager := pkg.NewPackageManager(pType)

		newPkg, err := askPackageName(pkgManager)
		if err != nil {
			return err
		}
		newPkg.Tags = askTags()

		store, err := config.Store.Read()
		if err != nil {
			return err
		}
		if store.Packages.Exists(newPkg) {
			return fmt.Errorf("This has already added.")
		}

		store.Packages.Add(newPkg)
		store.Save(config.Store.ReadFile())

		fmt.Println("Installed:", newPkg.Name)

		return nil
	})
}

func askPackageManager() (string, error) {
	pType := util.Ask("which package")
	if !pkg.ExistsManager(pType) {
		return "", fmt.Errorf("Invalid package type:", pType)
	}
	return pType, nil
}

func askPackageName(pkgManager pkg.PackageManager) (*pkg.Package, error) {
	pName := util.Ask("package name")
	if !pkgManager.Exists(pName) {
		return nil, fmt.Errorf("Not found package:", pName)
	}

	newPkg := &pkg.Package{
		Manager: pkgManager.Name(),
		Name:    pName,
	}

	return newPkg, nil
}

func askTags() []string {
	tagsStr := util.Ask("tags(split by space)")
	return strings.Split(tagsStr, " ")
}
