package cmds

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/otiai10/copy"

	"gopkg.in/src-d/go-git.v4"

	"github.com/itskass/gover/conf"
	"github.com/itskass/gover/repo"

	"gopkg.in/urfave/cli.v1"
)

// Install a version of golang
var Install = cli.Command{
	Name:      "install",
	Usage:     "Install a version of Golang",
	ArgsUsage: "install [version]",
	Action:    install,
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "use", Usage: "Sets the newly install version as current version"},
		cli.BoolFlag{Name: "skip-setup", Usage: "Used during to skip setup check", Hidden: true},
	},
}

func install(c *cli.Context) error {
	ver := formatVer(c.Args().Get(0))
	fmt.Printf("[Attepting to install version %s]\n", ver)

	if !c.Bool("skip-setup") {
		if err := requireSetup(); err != nil {
			return err
		}
	}

	// Open the source repo
	fmt.Println("- Opening Repo")
	r, err := git.PlainOpen(conf.Config.SourcePath)
	if err != nil {
		fmt.Printf("--> FAIL: Couldn't open repo %s\n", conf.Config.SourcePath)
		fmt.Printf("--> TRY: `gover upgrade`\n")
		return err
	}

	// check tags
	fmt.Println("- Checking Version Tag")
	tags, err := repo.Tags(r)
	if err != nil {
		return err
	}
	// Check version tag exists
	if !containsTag(ver, tags) {
		fmt.Printf("--> FAIL: unknown version %s\n", ver)
		return fmt.Errorf("bad version tag")
	}

	// copy source to version folder
	verPath := conf.Config.VerPath + "/" + ver
	fmt.Printf("- Copying Source to %s\n", verPath)
	conf.Spin.Start()
	os.RemoveAll(verPath)
	if err := copy.Copy(conf.Config.SourcePath, verPath); err != nil {
		fmt.Printf("--> FAIL: failed to copy source to version path\n")
		return err
	}
	conf.Spin.Stop()

	// open cloned repo
	fmt.Println("- Opening Repo")
	r, err = git.PlainOpen(verPath)
	if err != nil {
		fmt.Printf("--> FAIL: Couldn't open repo %s\n", conf.Config.SourcePath)
		fmt.Printf("--> TRY: `gover upgrade`\n")
		return err
	}

	// Get Working tree
	fmt.Println("- Getting WorkTree")
	tree, err := r.Worktree()
	if err != nil {
		fmt.Printf("--> FAIL: couldn't get work tree")
		return err
	}

	// checkout working tree
	fmt.Println("- Checking out tag")
	if err := repo.Checkout(tree, ver); err != nil {
		fmt.Printf("--> FAIL: failed to checkout tag/%s\n", ver)
		return err
	}

	// run make.bash
	fmt.Printf("- Running make.bash")
	conf.Spin.Start()
	if err := makeBashCommand(verPath).Run(); err != nil {
		fmt.Printf("--> FAIL: failed to run make.bash\n")
		return err
	}
	conf.Spin.Stop()

	// Installation is complete
	fmt.Printf("Installed: %s\n", ver)

	// Sets the newly install version as current version if the
	// `--use` flag is present.
	if c.Bool("use") {
		return use(c)
	}

	return nil
}

func formatVer(ver string) string {
	if !strings.HasPrefix("go", ver) {
		return "go" + ver
	}
	return ver
}

func containsTag(tag string, tags []string) bool {
	for _, t := range tags {
		if t == "refs/tags/"+tag {
			return true
		}
	}
	return false
}

func makeBashCommand(verPath string) *exec.Cmd {
	// set temporary goroot to version
	oldGOROOT := os.Getenv("GOROOT")
	defer os.Setenv("GOROOT", oldGOROOT)
	os.Setenv("GOROOT", verPath)
	// create command
	cmd := exec.Command(verPath + "/src/make.bash")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Dir = verPath + "/src"
	return cmd
}
