package repo

import (
	"strings"

	"github.com/itskass/gover/conf"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Tags returns all fetched tags that start with
// `go`. (i.e. ignores weekly tags)
func Tags(repo *git.Repository) ([]string, error) {
	conf.Spin.Start()
	defer conf.Spin.Stop()
	tagList := []string{}

	// get all tags
	tagrefs, err := repo.Tags()
	if err != nil {
		return nil, err
	}

	// interate tags and add valid tags to taglist
	tagrefs.ForEach(func(t *plumbing.Reference) error {
		tag := t.Name().String()
		if strings.HasPrefix(tag, "refs/tags/go") {
			tagList = append(tagList, tag)
		}
		return nil
	})

	return tagList, nil
}
