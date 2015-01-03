package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

// Download a file
func DownloadFile(url string, filename string, o *Output, force bool) error {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		o.ErrText("Could not download " + url)
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		o.ErrText("Could not dump body")
		os.Exit(1)
	}

	// Check if the file exists (and that force is not enabled)
	if _, err := os.Stat(filename); err == nil && (!force) {
		o.ErrText(filename + " already exists. Use -f to overwrite.")
		os.Exit(1)
	}

	// Check if the file contains "<html"
	if bytes.Contains(b, []byte("<html")) {
		return errors.New("Got html in return")
	}

	err = ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		o.ErrText("Could not write to " + filename + "!")
		os.Exit(1)
	}
	return err
}
