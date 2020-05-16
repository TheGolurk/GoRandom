package main

import (
	"os"
	"fmt"
	"unicode"
)

func main() {
	var arr []string
	for i := 1; i < len(os.Args); i++ {
		arr = append(arr, os.Args[i])
	}

	stringContent := ""
	for _, k := range arr {
		title := ""
		for i, letter := range k {
			if i < 1 {
				letter = unicode.ToUpper(letter)
			}
			title += string(letter)
		}
		stringContent += fmt.Sprintf("// %s ...\n", title)
		stringContent += fmt.Sprintf("type %s struct {\n\n}\n\n", title)
	}

	fmt.Println(stringContent)
}
