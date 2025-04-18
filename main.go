package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/xyproto/textoutput"
)

const versionString = "aget 1.4.0"

func run(o *textoutput.TextOutput, commandString string) error {
	var stdoutBuf, stderrBuf bytes.Buffer
	o.Println("<green>" + commandString + "</green>")
	words := strings.Fields(commandString)
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

func gitClone(o *textoutput.TextOutput, packageName string) error {
	sshURL := "ssh://aur@aur.archlinux.org/" + packageName + ".git"

	// Try SSH first
	if err := run(o, "git clone "+sshURL); err == nil {
		return nil
	}

	// Fallback to HTTPS by replacing scheme and removing user
	httpsURL := strings.Replace(sshURL, "ssh://aur@", "https://", 1)
	o.Println("<yellow>Falling back to HTTPS...</yellow>")
	return run(o, "git clone "+httpsURL)
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
				if _, statErr := os.Stat(packageName); statErr == nil {
					o.Print("<darkred>Directory already exists:</darkred> ")
					o.Println("<yellow>" + packageName + "</yellow>")
					continue
				}

				if cloneErr := gitClone(o, packageName); cloneErr != nil {
					err = cloneErr
					continue
				}

				// cd packageName
				o.Println("<green>cd " + packageName + "</green>")
				if chdirErr := os.Chdir(packageName); chdirErr != nil {
					o.Printf("<red>%s</red>\n", chdirErr)
					err = chdirErr
					continue
				}

				// switch to the master branch, in case the default branch name is ie. "main"
				if switchErr := run(o, "git switch -C master"); switchErr != nil {
					err = switchErr
					continue
				}
			}
			return err
		},
	}).Run(os.Args); appErr != nil {
		o.ErrExit(appErr.Error())
	}
}
