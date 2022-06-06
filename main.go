package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/xyproto/textoutput"
)

const versionString = "aget 1.3.1"

func run(o *textoutput.TextOutput, commandString string) error {
	var stdoutBuf, stderrBuf bytes.Buffer
	o.Println("<green>" + commandString + "</green>")
	words := strings.Split(commandString, " ")
	cmd := exec.Command(words[0], words[1:]...)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	if err := cmd.Start(); err != nil {
		o.Printf("<yellow>%s</yellow>\n", err)
		o.Printf("<yellow>%s</yellow>\n", stdoutBuf.String())
		o.Printf("<red>%s</red>\n", stderrBuf.String())
		return err
	}
	if err := cmd.Wait(); err != nil {
		o.Printf("<yellow>%s</yellow>\n", err)
		o.Printf("<yellow>%s</yellow>\n", stdoutBuf.String())
		o.Printf("<red>%s</red>\n", stderrBuf.String())
		return err
	}
	return nil
}

func main() {
	o := textoutput.New()
	if appErr := (&cli.App{
		Name:  "aget",
		Usage: "clone AUR packages with git",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "silent", Aliases: []string{"s"}},
			&cli.BoolFlag{Name: "https", Aliases: []string{"http", "web"}},
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

			// Check if any arguments are given
			if c.NArg() == 0 {
				o.ErrExit("Please supply a package name as an argument")
			}

			// Interpret the arguments as package names
			packageNames := c.Args().Slice()

			// Treat all arguments as AUR packages that should be cloned
			var err error
			for _, packageName := range packageNames {
				if _, err := os.Stat(packageName); err == nil {
					o.Print("<darkred>Directory already exists:</darkred> ")
					o.Println("<yellow>" + packageName + "</yellow>")
					continue
				}
				protocol := "ssh://"
				if c.Bool("https") {
					protocol = "https://"
				}
				url := fmt.Sprintf("%saur@aur.archlinux.org/%s.git", protocol, packageName)

				// git clone
				if err := run(o, "git clone "+url); err != nil {
					continue
				}

				// cd packageName
				o.Println("<green>cd " + packageName + "</green>")
				if err := os.Chdir(packageName); err != nil {
					o.Printf("<red>%s</red>\n", err)
					continue
				}

				// switch to the master branch, in case the default branch name is ie. "main"
				if err := run(o, "git switch -C master"); err != nil {
					continue
				}

			}
			return err
		},
	}).Run(os.Args); appErr != nil {
		o.ErrExit(appErr.Error())
	}

}
