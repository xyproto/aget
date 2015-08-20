package main

import (
	"github.com/xyproto/term"
	"os"
	"os/exec"
)

const (
	version = "0.4"
)

func isFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	return fileInfo.Mode().IsRegular(), err
}

func main() {
	o := term.NewTextOutput(true, true)

	// TODO: use the codegangsta/cli package

	if len(os.Args) < 2 {
		o.Println(o.LightBlue("aurtic " + version))
		o.ErrExit("Supply a package name as an argument")
	}

	pkg := os.Args[1]

	url := "ssh://aur@aur.archlinux.org/" + pkg + ".git"
	o.Println("git clone " + url)
	if _, err := exec.Command("git", "clone", url).Output(); err != nil {
		o.ErrExit("Could not clone " + pkg + ": " + err.Error())
	}

}
