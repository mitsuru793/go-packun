package commands

import (
	"os"

	"github.com/mitsuru793/go-packun/config"
	"github.com/mitsuru793/go-packun/util"
	"github.com/urfave/cli"
)

type subAction = func(c *cli.Context, configFile *config.Config) error

func action(c *cli.Context, subAction subAction) error {
	path := util.ExpandHome(c.GlobalString("config"))
	if path == "" {
		path = "~/.go-packun/config.toml"
	}
	path = util.ExpandHome(path)

	if util.Exists(path) {
		cfgFile, _ := config.Load(path)
		return subAction(c, cfgFile)
	} else {
		f, _ := os.Open(path)
		cfgFile := config.NewConfig()
		cfgFile.Save(f)
		return subAction(c, cfgFile)
	}
}
