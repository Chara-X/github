package github

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Registry struct{ Path string }

func (r *Registry) Push(registry, branch string) {
	filepath.Walk(r.Path, func(path string, info fs.FileInfo, err error) error {
		if isRepo(path) {
			var cmd = exec.Command("sh", "-c", fmt.Sprintln("git add --all; git commit --message Up; git push", registry+"/"+info.Name()+".git", "HEAD:"+branch))
			cmd.Dir = path
			var out, _ = cmd.Output()
			log.Println(path + "\n" + cmd.String() + "\n" + string(out))
		}
		return nil
	})
}
func (r *Registry) Pull(registry, branch string) {
	filepath.Walk(r.Path, func(path string, info fs.FileInfo, err error) error {
		if isRepo(path) {
			var cmd = exec.Command("sh", "-c", fmt.Sprintln("git add --all; git commit --message Up; git pull", registry+"/"+info.Name()+".git", branch))
			cmd.Dir = path
			var out, _ = cmd.Output()
			log.Println(path + "\n" + cmd.String() + "\n" + string(out))
		}
		return nil
	})
}
func isRepo(name string) bool {
	var entries, _ = os.ReadDir(name)
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() == ".git" {
			return true
		}
	}
	return false
}
