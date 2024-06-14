package doc_manager

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ZXSQ1/devdocs-tui/files"
	"github.com/ZXSQ1/devdocs-tui/utils"
)

// Errors

var ErrNotFetched error = fmt.Errorf("documentation entries not fetched")
var ErrNotFiltered error = fmt.Errorf("documentation entries not filtered")

/////

type DocsManager struct {
	languageName string
	docFile      string
}

/*
description: gets an instance of the DocsManager type
arguments:

	name: the name of the language

return: the DocsManager object with the language name
*/
func GetDocsManager(languageName string) DocsManager {
	var home = utils.GetEnvironVar("HOME")
	var docDir = home + "/.cache/devdocs-tui"
	var docFile = docDir + "/" + languageName

	if !files.IsExists(docDir) {
		os.MkdirAll(docDir, 0644)
	}

	return DocsManager{
		languageName: languageName,
		docFile:      docFile,
	}
}

/*
description: gets the documentation entries of the language
arguments: uses the fields in the DocsManager structure
return: a string containing the unfiltered documentation entries; stored in the DocsManager file
*/
func (docManager *DocsManager) FetchDocs() error {
	getDocsCMD := exec.Command("dedoc", "search", docManager.languageName)
	getDocsCMD.Stderr = os.Stderr
	getDocsCMD.Stdin = os.Stdin

	out, err := getDocsCMD.Output()

	if err != nil {
		return err
	}

	files.WriteFile(docManager.docFile, out)

	return nil
}

/*
description: filters the language documentation
arguments: uses the fields in the DocsManager structure
return: the filtered string documentation; stored in the DocsManager structure
*/
func (docManager *DocsManager) FilterDocs() {
	if docManager.isFiltered || !docManager.isFetched {
		return
	}

	unfilteredDocs := strings.ReplaceAll(docManager.docs, "\t", " ")

	result := ""
	parent := ""
	for _, line := range strings.Split(unfilteredDocs, "\n") {
		if !strings.HasPrefix(line, " ") {
			continue
		}

		words := strings.Split(line, " ")
		entry := words[len(words)-1]

		if strings.HasPrefix(entry, "#") {
			result += parent + entry + "\n"
		} else {
			parent = entry
			result += parent + "\n"
		}
	}

	docManager.docs = result
	docManager.isFiltered = true
}

/*
description: allows the user to choose docs
arguments: the fields in the DocsManager structure
return: the chosen doc is returned and stored in the DocsManager structure
*/
func (docManager *DocsManager) ChooseDocs() {
	if !docManager.isFiltered || !docManager.isFetched {
		return
	}

	cmd := exec.Command("bash", "-c", "fzf > .tmp && echo $(cat .tmp) && rm .tmp")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("ChooseDocs: error choosing documentation entry")
	}

	out, _ := cmd.Output()
	docManager.chosenDoc = string(out)
}
