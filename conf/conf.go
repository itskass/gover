package conf

import (
	"github.com/briandowns/spinner"
	"github.com/mitchellh/go-homedir"
)

type config struct {
	// SourceURl the url of the golang repo.
	SourceURL string `json:"source-url"`
	// RootPath - the folder root folder where all
	// gover directories are.
	RootPath string `json:"root-path"`
	// VerPath - where installed version binaries
	// are.
	VerPath string `json:"bin-path"`
	// SourcePath - where the main go repo is stored
	SourcePath string `json:"source-path"`

	// Branches Maps version numbers to branches
	Branches map[string]string `json:"branches"`
}

var Config *config
var Spin *spinner.Spinner

func SetDefault() {

	Config = &config{
		SourceURL: "https://github.com/golang/go",
	}

	// expand directories
	Config.RootPath, _ = homedir.Expand("~/.gover")
	Config.VerPath, _ = homedir.Expand("~/.gover/v")
	Config.SourcePath, _ = homedir.Expand("~/.gover/source")
}
