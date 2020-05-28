package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		err := tree(arg)
		if err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
	}
}

func tree(root string) error {
	err := filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.Name()[0] == '.' {
			return filepath.SkipDir
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return fmt.Errorf("Could not rel(%s, %s): %v", root, path, err)
		}
		fmt.Println(rel)

		fmt.Println(fi.Name())
		return nil
	})
	return err
}
