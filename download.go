package main

import (
	"bytes"
	"errors"
	"github.com/xyproto/textgui"
	"io/ioutil"
	"net/http"
	"os"
)

/*DownloadFile can download a file.
 *
 * o is for colored text
 * force is for allowing to overwrite the file when downloading
 * htmlcheck is for checking if the downloaded file contains "<html" or not
 */
func DownloadFile(url, filename string, o *textgui.TextOutput, force, htmlcheck bool) error {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		o.ErrExit("Could not download " + url)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		o.ErrExit("Could not dump body")
	}

	// Check if the file exists (and that force is not enabled)
	if _, err := os.Stat(filename); err == nil && (!force) {
		o.ErrExit(filename + " already exists. Use -f to overwrite.")
	}

	// Check if the file contains "<html"
	if htmlcheck && bytes.Contains(b, []byte("<html")) {
		return errors.New("Got html in return")
	}

	err = ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		o.ErrExit("Could not write to " + filename + "!")
	}
	return err
}
