package httprequests

import (
	"GoTranslator/common"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

// HTTPGetRequest makes a get request
func HTTPGetRequest(url string) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

	}

	//req.Header.Add("", ``)
	resp, err := client.Do(req)

	fmt.Println(resp.Body)

	if resp.StatusCode == http.StatusOK {
		convertedJSON := convertResponse(resp)

		fmt.Println(convertedJSON.Data.Title)
	}

}

func HTTPPostRequest(url string, input string, language string) {

	client := &http.Client{}

	var something = []byte(input)

	req, err := http.NewRequest("POST", url+"&to=en", bytes.NewBuffer(something))

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
