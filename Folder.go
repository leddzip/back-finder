package main

import (
	"errors"
	"fmt"
	"os"
	"path"
)

type Folder struct {
	Path string
}

// ContainsFile return a boolean telling if the given filename exist
// in the folder or an error if it is impossible to infer if
// the file exist or not.
func (f *Folder) ContainsFile(fileName string) (bool, error) {
	filePathCandidate := path.Join(f.Path, fileName)
	_, err := os.Stat(filePathCandidate)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, errors.New(fmt.Sprintf("unable to determine if the file %s exist or not", filePathCandidate))
	}
}

// NavigateToParent update the current Folder one path above ("..")
// it then returns the new path of the folder, as a string
func (f *Folder) NavigateToParent() string {
	f.Path = path.Clean(path.Join(f.Path, ".."))
	return f.Path
}
