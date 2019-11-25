package cmds

import (
	"fmt"
	"os"

	"github.com/itskass/gover/conf"
	"gopkg.in/urfave/cli.v1"
)

// Setup sets up up gover and installs Golang version 1.13 as the base base
var Setup = cli.Command{
	Name:   "setup",
	Usage:  "sets up gover and installs Golang version 1.13 as the base base",
	Action: setup,
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "force", Usage: "force setup even if paths already exist"},
	},
}

func setup(c *cli.Context) error {
	fmt.Println("[Running Setup]")

	if !c.Bool("force") {
		if pathExists(conf.Config.RootPath + "/goroot") {
			fmt.Println("--> FAIL: goroot already setup")
			fmt.Println("--> TRY: `gover setup --force`")
			return nil
		}
	}

	// create root directory
	fmt.Println("- SETUP: creating ", conf.Config.RootPath)
	if _, err := os.Stat(conf.Config.RootPath); os.IsNotExist(err) {
		if err := os.MkdirAll(conf.Config.RootPath, 0700); err != nil {
			return err
		}
		if err := os.MkdirAll(conf.Config.VerPath, 0700); err != nil {
			return err
		}
	}

	// upgrade which downloads the go source code from the source
	// path
	fmt.Println("- SETUP: fetching Golang source code...")
	if err := c.App.Run([]string{os.Args[0], "fetch"}); err != nil {
		return err
	}

	// Install go version 1.13
	fmt.Println("- SETUP: installing go version 1.13 ...")
	if err := c.App.Run([]string{os.Args[0], "install", "1.13", "--use", "--skip-setup"}); err != nil {
		return nil
	}

	return nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func requireSetup() error {
	if !pathExists(conf.Config.RootPath + "/goroot") {
		fmt.Println("--> FAIL: Setup Required")
		fmt.Println("--> TRY: `gover setup --force`")
		return fmt.Errorf("Setup Required")
	}
	return nil
}
