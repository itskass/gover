package repo

import (
	"github.com/itskass/gover/conf"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Checkout a tag from the worktree
func Checkout(tree *git.Worktree, tag string) error {
	conf.Spin.Start()
	defer conf.Spin.Stop()
	return tree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewTagReferenceName(tag),
		Create: false,
		Force:  true,
	})
}
