package cmds

import (
	"os"
	"path/filepath"

	"github.com/itskass/gover/conf"
)

func setCurrent(ver string) error {
	os.Remove(conf.Config.RootPath + "/goroot")
	verPath := conf.Config.VerPath + "/" + ver
	return os.Symlink(verPath, conf.Config.RootPath+"/goroot")
}

func getCurrent() (string, error) {
	return filepath.EvalSymlinks(conf.Config.RootPath + "/goroot")
}
