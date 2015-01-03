package main

import (
	"os"
	"os/exec"
)

const (
	ext  = ".tar.gz"
	repo = "https://aur.archlinux.org/packages/"
)

func main() {
	o := NewOutput(true, true)

	if len(os.Args) < 2 {
		o.ErrText("Need a package name as the first argument")
		os.Exit(1)
	}

	// Parse the arguments
	force := false
	pkg := os.Args[1]
	if len(os.Args) >= 3 {
		if os.Args[1] == "-f" {
			force = true
			pkg = os.Args[2]
		}
		if os.Args[2] == "-f" {
			force = true
			pkg = os.Args[1]
		}
	}

	// Download the file
	url := repo + pkg[:2] + "/" + pkg + "/" + pkg + ext
	if err := DownloadFile(url, pkg+ext, o, force); err != nil {
		o.ErrText("Could not download " + pkg + " from AUR.")
		os.Exit(1)
	}

	// Check if the directory exists (and that force is not enabled)
	if _, err := os.Stat(pkg); err == nil && (!force) {
		o.ErrText(pkg + " already exists. Use -f to overwrite.")
		os.Exit(1)
	}

	// Extract the file
	cmd := exec.Command("tar", "zxf", pkg+ext)
	_, err := cmd.Output()
	if err != nil {
		o.ErrText("Could not extract " + pkg + ext)
		os.Exit(1)
	}

	// Remove the file
	cmd = exec.Command("rm", pkg+ext)
	_, err = cmd.Output()
	if err != nil {
		o.ErrText("Could not remove " + pkg + ext)
		os.Exit(1)
	}

}
