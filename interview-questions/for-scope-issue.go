package main

import (
	"fmt"
	"os"
)

func main() {
	var rmdirs []func()
	for _, dir := range tempDirs() {
		dir := dir
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir) // NOTE: incorrect!
		})
		fmt.Println(dir)
	}

	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}

func tempDirs() []string {
	return []string{"tempDir1", "tempDir2", "tempDir3", "tempDir4"}
}
