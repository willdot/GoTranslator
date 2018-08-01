package main

import (
	"GoTranslator/translate"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var input string
	quit := false

	for quit != true {

		fmt.Println("Enter something to translate or type q to quit")
		input = getUserInput()

		if input != "q" {

			var language string
			languagePointer := &language

			languageCheck := checkLanguage(languagePointer)
			fmt.Println(languageCheck)

			if quit == false && languageCheck == true {
				translate.Translate(input, language)
			} else {
				quit = true
			}

		} else {
			quit = true
		}
	}

	fmt.Println("Thankyou for trying translator")
}

func getUserInput() string {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// remove the new line
	text = strings.TrimSuffix(text, "\r\n")
	text = strings.TrimSuffix(text, "\n")

	return text
}

func checkLanguage(language *string) bool {

	success := false

	for success == false {
		fmt.Println("Please enter a language you wish to translate into (en, fr, de, it) or q to quit")

		*language = getUserInput()

		if *language == "q" {
			return false
		}

		if *language != "en" && *language != "fr" && *language != "de" && *language != "it" {
			fmt.Println("Incorrect language selection")
		} else {
			success = true
		}
	}

	return true
}
