package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/willdot/GoTranslator/translate"
)

func main() {

	var input string
	quit := false

	// Loop to keep application running until user selects to quit
	for quit != true {

		fmt.Println("Enter something to translate or type q to quit")
		input = getUserInput()

		// User has chosen to quit, so break out the loop and exit application
		if input == "q" {
			break
		}

		var language string
		languagePointer := &language

		languageCheck := checkLanguage(languagePointer)

		// Check that the selected language is valid. If false, the user selected to quit
		if languageCheck == true {
			translate.Translate(input, language)
		} else {
			// User has chosen to quit, so break out the loop and exit application
			break
		}

	}

	fmt.Println("Thankyou for trying translator")
}

func checkLanguage(language *string) bool {

	success := false

	// Loop through until either the user selects to quit or they enter a valid language code
	for success == false {
		fmt.Println("Please enter a language you wish to translate into (en, fr, de, it) or q to quit")

		*language = getUserInput()

		// User has chosen to quit, so break out the loop to return false
		if *language == "q" {
			success = false
			break
		}

		// Make sure the user has entered a valid language code
		if *language != "en" && *language != "fr" && *language != "de" && *language != "it" {
			fmt.Println("Incorrect language selection")
		} else {
			success = true
		}
	}

	return success
}

func getUserInput() string {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// remove the new line
	text = strings.TrimSuffix(text, "\r\n")
	text = strings.TrimSuffix(text, "\n")

	return text
}
