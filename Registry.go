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
	Path   string
	Ignore []string
}

func (r *Registry) Push(registry, branch string) {
	// var wg = sync.WaitGroup{}
	filepath.Walk(r.Path, func(path string, info fs.FileInfo, err error) error {
		if isRepo(path) && !slices.Contains(r.Ignore, info.Name()) {
			// wg.Add(1)
			// go func() {
			var cmd = exec.Command("sh", "-c", fmt.Sprintln("git add --all; git commit --message Up; git push", registry+"/"+info.Name()+".git", "HEAD:"+branch))
			cmd.Dir = path
			var out, _ = cmd.Output()
			log.Println(path + "\n" + cmd.String() + "\n" + string(out))
			// wg.Done()
			// }()
		}
		return nil
	})
	// wg.Wait()
}
func (r *Registry) Pull(registry, branch string) {
	var wg = sync.WaitGroup{}
	filepath.Walk(r.Path, func(path string, info fs.FileInfo, err error) error {
		if isRepo(path) && !slices.Contains(r.Ignore, info.Name()) {
			wg.Add(1)
			go func() {
				var cmd = exec.Command("sh", "-c", fmt.Sprintln("git add --all; git commit --message Up; git pull", registry+"/"+info.Name()+".git", branch))
				cmd.Dir = path
				var out, _ = cmd.Output()
				log.Println(path + "\n" + cmd.String() + "\n" + string(out))
				wg.Done()
			}()
		}
		return nil
	})
	wg.Wait()
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
