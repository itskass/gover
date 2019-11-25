package main

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/itskass/gover/conf"

	"github.com/itskass/gover/cmds"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	conf.SetDefault()

	// Define CLI Application
	app := cli.NewApp()
	app.Name = "gover"
	app.Usage = "Install and switch between multiple versions of Golang"
	app.Author = "itskass <itskass94@gmail.com>"
	app.Version = "0.5"
	app.HideVersion = true

	// Register commands
	app.Commands = []cli.Command{
		cmds.Use,
		cmds.Install,
		cmds.Fetch,
		cmds.List,
		cmds.Setup,
	}
	// create spinner
	conf.Spin = spinner.New(spinner.CharSets[28], 100*time.Millisecond)
	conf.Spin.Prefix = "working: "

	// Check for first run
	firstRun(app)

	// Run CLI Application
	if err := app.Run(os.Args); err != nil {
		fmt.Println("\n!!! ERROR:", err)
		os.Exit(1)
	}
}

func firstRun(app *cli.App) {
	if _, err := os.Stat(conf.Config.RootPath); os.IsNotExist(err) {
		fmt.Println("[FIRST RUN DETECTED]")
		app.Run([]string{os.Args[0], "setup", "--force=true"})
	}
}
