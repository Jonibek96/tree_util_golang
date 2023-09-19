package main

import (
	"os"
	"tree_util_golang/tree"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	path := "testdata"
	err := tree.DirTree(out, path, printFiles)
	if err != nil {
		panic(err)
	}
}
