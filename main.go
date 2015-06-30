package main

import (
	"github.com/xyproto/term"
	"os"
	"os/exec"
	"strings"
	//"path/filepath"
)

const (
	ext     = ".tar.gz"
	repo    = "https://aur.archlinux.org/packages/"
	version = "0.3"
)

func isFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	return fileInfo.Mode().IsRegular(), err
}

func main() {
	o := term.NewTextOutput(true, true)

	if len(os.Args) < 2 {
		o.Println(o.LightBlue("aurtic " + version))
		o.ErrExit("Supply a package name as an argument")
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

	// If the name ends with ".git", clone from AUR4
	if strings.HasSuffix(pkg, ".git") {
		url := "ssh+git://aur@aur4.archlinux.org/" + pkg + "/"
		if _, err := exec.Command("git", "clone", url).Output(); err != nil {
			o.ErrExit("Could not clone " + pkg)
		}
		return
	}

	url := repo + pkg[:2] + "/" + pkg + "/" + pkg + ext

	// Download the file
	if err := DownloadFile(url, pkg+ext, o, force, true); err != nil {
		o.ErrExit("Could not download " + pkg + " from AUR.")
	}

	// Check if the directory exists (and that force is not enabled)
	if _, err := os.Stat(pkg); err == nil {
		if force {
			// Remove the directory
			if _, err := exec.Command("rm", "-rf", pkg).Output(); err != nil {
				o.ErrExit("Could not remove: " + pkg)
			}
		} else {
			o.ErrExit(pkg + " already exists. Use -f to overwrite.")
		}
	}

	// Extract the file
	if _, err := exec.Command("tar", "zxf", pkg+ext).Output(); err != nil {
		o.ErrExit("Could not extract " + pkg + ext)
	}

	// Setting premissions for the files, after extracting
	//matches, err := filepath.Glob(pkg + "/*")
	//if err != nil {
	//	o.ErrExit("Could not glob")
	//}
	//for _, filename := range matches {
	//	if file, err := isFile(filename); file && (err != nil) {
	//		// Set the permissions to 644
	//		o.Println("changing perms for " + pkg + "/" + filename)
	//		if _, err := exec.Command("chmod", "644", pkg+"/"+filename).Output(); err != nil {
	//			o.ErrExit("Could not set permissions for " + pkg+"/"+filename)
	//		}
	//	}
	//}

	// Remove the file
	if _, err := exec.Command("rm", pkg+ext).Output(); err != nil {
		o.ErrExit("Could not remove " + pkg + ext)
	}

}
