package cmds

import (
	"fmt"
	"os"

	"github.com/itskass/gover/conf"
	"gopkg.in/urfave/cli.v1"
)

// Use will assign a current version of Golang to use
var Use = cli.Command{
	Name:      "use",
	Usage:     "Assign an installed version to use",
	ArgsUsage: "use [version]",
	Action:    use,
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "skip-setup", Usage: "Used during to skip setup hceck", Hidden: true},
	},
}

func use(c *cli.Context) error {
	ver := formatVer(c.Args().Get(0))
	verPath := conf.Config.VerPath + "/" + ver
	fmt.Printf("[Setting Current Version to %s]\n", ver)

	// check setup
	if !c.Bool("skip-setup") {
		if err := requireSetup(); err != nil {
			return err
		}
	}

	// check binary exists
	if _, err := os.Stat(verPath + "/bin/go"); os.IsNotExist(err) {
		fmt.Println("--> FAIL: version not installed")
		fmt.Printf("--> TRY: `gover install %s`", ver)
		return err
	}

	return setCurrent(ver)
}
