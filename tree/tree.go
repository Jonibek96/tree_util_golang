package tree

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"sort"
)

func DirTree(out io.Writer, path string, printFiles bool) error {
	err := dirTreeHelper(out, path, printFiles, "")
	if err != nil {
		return err
	}
	return nil
}

func dirTreeHelper(out io.Writer, path string, printFiles bool, prefix string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	if !printFiles {
		var newFiles []fs.FileInfo
		for _, file := range files {
			if file.IsDir() {
				newFiles = append(newFiles, file)
			}
		}
		files = nil
		files = newFiles
	}

	for i, file := range files {
		newPrefix := prefix
		if !printFiles && !file.IsDir() {
			continue
		}
		isLast := i == len(files)-1
		if !isLast {
			fmt.Fprintf(out, "%s├───", prefix)
			newPrefix += "│\t"
		} else {
			fmt.Fprintf(out, "%s└───", prefix)
			newPrefix += "\t"
		}
		if i == len(files)-1 {
			newPrefix += ""

		}
		if file.IsDir() {
			fmt.Fprintf(out, "%s\n", file.Name())
			err := dirTreeHelper(out, filepath.Join(path, file.Name()), printFiles, newPrefix)
			if err != nil {
				return err
			}
		} else {
			if file.Size() == 0 {
				fmt.Fprintf(out, "%s (empty)\n", file.Name())
			} else {
				fmt.Fprintf(out, "%s (%vb)\n", file.Name(), file.Size())
			}
		}
	}

	return nil
}
