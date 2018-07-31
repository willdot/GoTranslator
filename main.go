package main

import (
	"GoTranslator/common"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type ResponseJson struct {
	Data   Todo `json:"data"`
	Status int  `json:"status"`
}

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	/*var input string

	for input != "q" {
		fmt.Println("Enter something to translate or type q to quit")
		input = getUserInput()

		if input != "q" {

			output := translate.Translate(input)

			fmt.Println("Translation: ", output)
		}
	}*/

	HTTPPostRequest("https://api.cognitive.microsofttranslator.com/translate?api-version=3.0", "bonjour", "en")

	fmt.Println("Thankyou for trying translator")
}

func getUserInput() string {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	//fmt.Println(text)

	// remove the new line
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func HTTPPostRequest(url string, input string, language string) {

	client := &http.Client{}

	url = url + "&to=en"

	//var a = bytes.NewBuffer(something)

	body := []byte(input)
	//req, err := http.Post("http://someurl.com", "body/type", bytes.NewBuffer(body))
	// Code to process response (written in Get request snippet) goes

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {

	}

	apiKey := common.GetAPIKey()

	req.Header.Add("Ocp-Apim-Subscription-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	fmt.Println(resp.Body)

	if resp.StatusCode == http.StatusOK {
		convertedJSON := convertResponse(resp)

		fmt.Println(convertedJSON.Data.Title)
	}
}

func convertResponse(resp *http.Response) *ResponseJson {

	decoder := json.NewDecoder(resp.Body)

	var result = new(ResponseJson)

	err := decoder.Decode(&result)

	if err != nil {
		fmt.Println(err)
	}
	return result
}
