package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var input = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Print("Put something: ")
	input.Scan()
	text := input.Text()

	fmt.Println(dictionary(text))

}

func dictionary(text string) string {

	var isWord = regexp.MustCompile(`^[A-Z]`)
	letters := make(map[string]string)
	word := " "
	temp := ""

	letters["A"] = ".-"
	letters["B"] = "-..."
	letters["C"] = "-.-."
	letters["D"] = "-.."
	letters["E"] = "."
	letters["F"] = "..-."
	letters["G"] = "--."
	letters["H"] = "...."
	letters["I"] = ".."
	letters["J"] = ".---"
	letters["K"] = "-.-"
	letters["L"] = ".-.."
	letters["M"] = "--"
	letters["N"] = "-."
	letters["O"] = "---"
	letters["P"] = ".--."
	letters["Q"] = "--.-"
	letters["R"] = ".-."
	letters["S"] = "..."
	letters["T"] = "-"
	letters["U"] = "..-"
	letters["V"] = "...-"
	letters["W"] = ".--"
	letters["X"] = "-..-"
	letters["Y"] = "-.--"
	letters["Z"] = "--.."

	if isWord.MatchString(strings.ToUpper(text)) {

		for _, value := range text {
			for index := range letters {
				if index == strings.ToUpper(string(value)) {
					word += " " + letters[string(index)]
				}
			}
			if string(value) == " " {
				word += " "
			}
		}

	} else {

		for _, value := range text {
			if value == 32 {
				word += " " + toEnglish(letters, temp)
				temp = ""
			} else {
				temp += string(value)
			}
		}
		word += toEnglish(letters, temp)

	}

	return word
}

func toEnglish(letters map[string]string, temp string) string {
	word := ""
	for index, morse := range letters {
		if morse == temp {
			word += index
		}
	}
	return word
}
