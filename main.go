package main

import (
	"bytes"
	"github.com/urfave/cli/v2"
	"github.com/xyproto/textoutput"
	"os"
	"os/exec"
)

const versionString = "aget 1.0.0"

func isFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.Mode().IsRegular(), nil
}

func main() {
	o := textoutput.New()
	if appErr := (&cli.App{
		Name:  "aget",
		Usage: "clone AUR packages with git",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "silent", Aliases: []string{"s"}},
			&cli.BoolFlag{Name: "version", Aliases: []string{"V"}},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("version") {
				o.Println(versionString)
				os.Exit(0)
			}
			// Check if text output should be disabled
			if c.Bool("silent") {
				o.Disable()
			}
			packageNames := []string{}
			// Check if any arguments are given
			if c.NArg() > 0 {
				packageNames = c.Args().Slice()
			} else {
				o.ErrExit("Supply a package name as an argument")
			}

			// Treat all arguments as AUR packages that should be cloned
			var err error
			for _, packageName := range packageNames {
				if _, err := os.Stat(packageName); err == nil {
					o.Print("<darkred>Directory already exists:</darkred> ")
					o.Println("<yellow>" + packageName + "</yellow>")
					continue
				}
				url := "ssh://aur@aur.archlinux.org/" + packageName + ".git"
				o.Println("<green>git clone " + url + "</green>")
				cmd := exec.Command("git", "clone", url)
				var stdoutBuf, stderrBuf bytes.Buffer
				cmd.Stdout = &stdoutBuf
				cmd.Stderr = &stderrBuf
				if err := cmd.Start(); err != nil {
					o.Printf("<yellow>%s</yellow>\n", err)
					o.Printf("<yellow>%s</yellow>\n", stdoutBuf.String())
					o.Printf("<red>%s</red>\n", stderrBuf.String())
					continue
				}
				if err := cmd.Wait(); err != nil {
					o.Printf("<yellow>%s</yellow>\n", err)
					o.Printf("<yellow>%s</yellow>\n", stdoutBuf.String())
					o.Printf("<red>%s</red>\n", stderrBuf.String())
					continue
				}
			}
			return err
		},
	}).Run(os.Args); appErr != nil {
		o.ErrExit(appErr.Error())
	}

}
