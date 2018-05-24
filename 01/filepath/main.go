package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	//Path to executable
	fmt.Println(ex)

	dir := filepath.Dir(ex)
	fmt.Println("Executable path: " + dir)

	symlinks, err := filepath.EvalSymlinks(dir)

	if err != nil {
		panic(err)
	}

	fmt.Println("Symlink evaluated:" + symlinks)
}
