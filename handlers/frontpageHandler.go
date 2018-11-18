package handlers

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"io/ioutil"
	"net/http"
)

// Frontpage is a webpage explaining api functionality
func Frontpage(w http.ResponseWriter, r *http.Request) {

	// Loads the README file.
	file, err := ioutil.ReadFile("README.md")
	if err != nil {
		logging("error, could not read the markdown file")
		fmt.Fprintln(w, "post to /dialogflow ")
	} else {
		// Converts markdown to byte stream.
		markdownData := []byte(file)

		// Converts the markdown file to html.
		html := markdown.ToHTML(markdownData, nil, nil)
		w.Write([]byte(html))

	}
}
