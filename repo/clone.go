package repo

import (
	"os"

	"github.com/itskass/gover/conf"
	"gopkg.in/src-d/go-git.v4"
)

// DeleteSourceRepo deletes any pre-existing source folder
func DeleteSourceRepo() error {
	conf.Spin.Start()
	defer conf.Spin.Stop()
	return os.RemoveAll(conf.Config.SourcePath)
}

// CloneSourceRepo clones the source repo from the path specified in the
// config.
func CloneSourceRepo() (*git.Repository, error) {
	conf.Spin.Start()
	defer conf.Spin.Stop()
	return git.PlainClone(conf.Config.SourcePath, false, &git.CloneOptions{
		URL:      conf.Config.SourceURL,
		Progress: os.Stdout,
	})
}
