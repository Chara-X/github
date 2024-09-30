package main

import (
	"fmt"
	"os"

	"github.com/Chara-X/github"
)

func main() {
	var reg = github.Registry{Exclude: os.Args[4:]}
	reg.Path, _ = os.Getwd()
	fmt.Println(reg.Path)
	switch os.Args[1] {
	case "push":
		reg.Push(os.Args[2], os.Args[3])
	case "pull":
		reg.Pull(os.Args[2], os.Args[3])
	}
}
