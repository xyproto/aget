package main

import (
	"os"
	"os/exec"
)

const (
	ext     = ".tar.gz"
	repo    = "https://aur.archlinux.org/packages/"
	version = "0.1"
)

func main() {
	o := NewOutput(true, true)

	if len(os.Args) < 2 {
		o.Println(o.LightBlue("aurtic " + version))
		o.Exit("Supply a package name as an argument")
	}

	// TODO: use a package for optparsing

	// Parse the arguments
	force := false
	pkg := os.Args[1]
	if len(os.Args) == 3 {
		if os.Args[1] == "-f" {
			force = true
			pkg = os.Args[2]
		} else if os.Args[2] == "-f" {
			force = true
			pkg = os.Args[1]
		}
	}

	// Download the file
	url := repo + pkg[:2] + "/" + pkg + "/" + pkg + ext
	if err := DownloadFile(url, pkg+ext, o, force, true); err != nil {
		o.Exit("Could not download " + pkg + " from AUR.")
	}

	// Check if the directory exists (and that force is not enabled)
	if _, err := os.Stat(pkg); err == nil && (!force) {
		o.Exit(pkg + " already exists. Use -f to overwrite.")
	}

	// Extract the file
	cmd := exec.Command("tar", "zxf", pkg+ext)
	if _, err := cmd.Output(); err != nil {
		o.Exit("Could not extract " + pkg + ext)
	}

	// Remove the file
	cmd = exec.Command("rm", pkg+ext)
	if _, err := cmd.Output(); err != nil {
		o.Exit("Could not remove " + pkg + ext)
	}

}
