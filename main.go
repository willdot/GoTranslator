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

	for input != "q" {
		fmt.Println("Enter something to translate or type q to quit")
		input = getUserInput()

		if input != "q" {

			output := translate.Translate(input)

			fmt.Println("Translation: ", output)
		}
	}

	fmt.Println("Thankyou for trying translator")
}

func makeHTTPRequest() {

}

func convertResponseToModel() {

}

func getUserInput() string {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	//fmt.Println(text)

	// remove the new line
	text = strings.Replace(text, "\n", "", -1)
	return text
}
