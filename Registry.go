package git

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

type Registry struct{ Path string }

func (r *Registry) Push(repository, branch string) error {
	return filepath.Walk(r.Path, func(path string, info fs.FileInfo, err error) error {
		if isRepo(path) {
			var cmd = exec.Command("sh", "-c", fmt.Sprintln("git add --all && git commit --message Up && git push", filepath.Join(repository, info.Name()+".git"), "HEAD:"+branch))
			cmd.Dir = r.Path
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		}
		return nil
	})
}
func (r *Registry) Pull(repository, branch string) error {
	return filepath.Walk(r.Path, func(path string, info fs.FileInfo, err error) error {
		if isRepo(path) {
			var cmd = exec.Command("sh", "-c", fmt.Sprintln("git add --all && git commit --message Up && git pull", filepath.Join(repository, info.Name()+".git"), branch))
			cmd.Dir = r.Path
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
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
