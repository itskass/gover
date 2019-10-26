package cmds

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/itskass/gover/repo"
	"gopkg.in/src-d/go-git.v4"

	"github.com/itskass/gover/conf"
	"github.com/urfave/cli"
)

// List will assign a current version of Golang to use
var List = cli.Command{
	Name:      "list",
	Usage:     "list all/installed versions",
	ArgsUsage: "use [version]",
	Subcommands: []cli.Command{
		// List all
		{
			Name:  "all",
			Usage: "list all avaliable versions",
			Action: func(c *cli.Context) error {
				fmt.Println("[Avaliable Versions]")
				return printTags()
			},
		},
		// list installed
		{
			Name:  "installed",
			Usage: "list all installed versions",
			Action: func(c *cli.Context) error {
				fmt.Println("[Installed Versions]")
				return printVerDir()
			},
		},
	},
}

func printTags() error {
	// get repo
	r, err := git.PlainOpen(conf.Config.SourcePath)
	if err != nil {
		fmt.Printf("--> FAIL: Couldn't open repo %s\n", conf.Config.SourcePath)
		fmt.Printf("--> TRY: `gover upgrade`\n")
		return err
	}
	// get tags
	conf.Spin.Start()
	tags, err := repo.Tags(r)
	if err != nil {
		return err
	}
	conf.Spin.Stop()
	// print all tags
	for _, ver := range tags {
		fmt.Printf("- %s\n", filepath.Base(ver))
	}
	return nil
}

func printVerDir() error {
	// get current version number
	curr, err := getCurrent()
	if err != nil {
		return err
	}
	curr = filepath.Base(curr)
	// get all direcotires in verpath
	files, err := ioutil.ReadDir(conf.Config.VerPath)
	if err != nil {
		return err
	}
	// print version, prefixing * to the current
	for _, file := range files {
		if curr == file.Name() {
			fmt.Printf("* %s\n", file.Name())
			continue
		}
		fmt.Printf("- %s\n", file.Name())
	}
	return nil
}
