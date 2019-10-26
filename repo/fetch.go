package repo

import (
	"os"

	"github.com/itskass/gover/conf"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
)

// FetchAll tags and branches
func FetchAll(remote *git.Remote) error {
	conf.Spin.Start()
	defer conf.Spin.Stop()
	return remote.Fetch(&git.FetchOptions{
		RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
		Progress: os.Stdout,
	})
}
