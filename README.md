# Gover

Gover is a command line tool for installing and managing multiple versions of Golang. Gover allows installation and
usage of Go1-Go13+ using the github tags; For supported* version see [Go Repo Tags](https://github.com/golang/go/tags).

*\*You may have trouble building older version of go.*

## Installation 
1. Get and install gover:
```bash
$ go get github.com/itskass/gover
```

2. Add gover to path:
```bash
# Gover must be prepended to ensure make sure the bash command checks the gover for
# golang first. If no version is set then it will fall back to your current version.
export PATH=~/.gover/goroot:$PATH
```

3. Run gover setup: 
```bash
$ gover setup 
```
Setup will clone the go repo, then install go1.13 and set it as the current version

## Quick start
To install and use another version of Go:
```bash
$ gover install 1.10.2 
$ gover use 1.10.2
```
or the short hand method:
```bash
$ gover install 1.9 --use
```

## Updating Gover
If a new version of go has been released, gover will need to fetch from the golang github repo to access the tag. 
To do this use the gover fetch command:
```bash
$ gover fetch
```

Contributions welcomed <3