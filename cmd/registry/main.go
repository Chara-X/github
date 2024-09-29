package main

import (
	"os"

	"github.com/Chara-X/git"
)

func main() {
	var workdir, _ = os.Getwd()
	var reg = git.Registry{Path: workdir}
	switch os.Args[1] {
	case "push":
		reg.Push(os.Args[2], os.Args[3])
	case "pull":
		reg.Pull(os.Args[2], os.Args[3])
	}
}
