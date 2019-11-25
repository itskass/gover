package cmds

import (
	"fmt"

	"github.com/itskass/gover/conf"

	"github.com/itskass/gover/repo"

	"gopkg.in/urfave/cli.v1"
)

// Fetch command pulls from the current Golang source repo
var Fetch = cli.Command{
	Name:   "fetch",
	Usage:  "fetch current Golang source repo",
	Action: fetch,
}

// fetch the latest Go source files from the source url
func fetch(c *cli.Context) error {
	fmt.Println("[Fetching latest Go Source files]")

	// delete pre-existing repo
	fmt.Printf("- Delete old source Folder: %s\n", conf.Config.SourcePath)
	if err := repo.DeleteSourceRepo(); err != nil {
		return err
	}

	// clone repo
	fmt.Printf("- Cloning from Git repo: %s\n", conf.Config.SourceURL)
	r, err := repo.CloneSourceRepo()
	if err != nil {
		fmt.Printf("--> FAIL: failed to clone %s\n", conf.Config.SourceURL)
		return err
	}

	// get remote origin
	remote, err := r.Remote("origin")
	if err != nil {
		fmt.Printf("--> FAIL: failed to get remote origin\n")
		return nil
	}

	// fetch all
	fmt.Printf("- Fetching Branches and Tags: \n")
	err = repo.FetchAll(remote)

	return err
}
