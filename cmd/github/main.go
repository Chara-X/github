package main

import (
	"os"

	"github.com/Chara-X/github"
)

func main() {
	var reg = github.Registry{}
	reg.Path, _ = os.Getwd()
	switch os.Args[1] {
	case "push":
		reg.Push(os.Args[2], os.Args[3])
	case "pull":
		reg.Pull(os.Args[2], os.Args[3])
	}
}
