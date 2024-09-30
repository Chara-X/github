package github

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"sync"
)

type Registry struct {
	Path    string
	Exclude []string
}

func (r *Registry) Push(registry, branch string) {
	var wg = sync.WaitGroup{}
	filepath.Walk(r.Path, func(path string, info fs.FileInfo, err error) error {
		if isRepo(path) && !slices.Contains(r.Exclude, info.Name()) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var cmd = exec.Command("sh", "-c", fmt.Sprintln("git add --all; git commit --message Up; git push", registry+"/"+info.Name()+".git", "HEAD:"+branch))
				cmd.Dir = path
				var out, _ = cmd.CombinedOutput()
				log.Println(path + "\n" + string(out) + "\n" + cmd.String())
			}()
		}
		return nil
	})
	wg.Wait()
}
func (r *Registry) Pull(registry, branch string) {
	filepath.Walk(r.Path, func(path string, info fs.FileInfo, err error) error {
		if isRepo(path) && !slices.Contains(r.Exclude, info.Name()) {
			log.Println(path)
			log.Println(fmt.Sprintln("git add --all; git commit --message Up; git pull", registry+"/"+info.Name()+".git", branch))
			var cmd = exec.Command("sh", "-c", fmt.Sprintln("git add --all; git commit --message Up; git pull", registry+"/"+info.Name()+".git", branch))
			cmd.Dir = path
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
