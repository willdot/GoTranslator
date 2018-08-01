package httprequests

import (
	"GoTranslator/common"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// AzureTranslationResponse is a model to map the Azure response
type AzureTranslationResponse struct {
	DetectedLan  DetectedLanguage `json:"detectedLanguage"`
	Translations []Translations   `json:"translations"`
}

// DetectedLanguage is a model to map the detected language Azure response
type DetectedLanguage struct {
	Language string  `json:"language"`
	Score    float32 `json:"score"`
}

// Translations is a model to map the translation result in the Azure response
type Translations struct {
	Text string `json:"text"`
	To   string `json:"to"`
}

// HTTPPostRequest makes a call to the Azure text translation service to get translations
func HTTPPostRequest(url string, input string, language string) *AzureTranslationResponse {

	apiKey := common.GetAPIKey()

	if apiKey == "" {
		fmt.Println("Error getting API key")
		return nil
	}

	client := &http.Client{}

	url = url + "&to=" + language

	body := strings.NewReader("[{\"Text\" : \"" + input + "\"}]")

	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		fmt.Printf("Error on request: %v\n", err)
		return nil
	}

	req.Header.Add("Ocp-Apim-Subscription-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error on request: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		convertedJSON := convertResponse(resp)

		return convertedJSON
	}

	return nil
}

func convertResponse(resp *http.Response) *AzureTranslationResponse {

	// The json response that Azure returns wraps the response model is an array, so this will remove the [ ] from the begining and end of the response json
	responseWithoutArray := removeArrayFromResponse(resp)

	decoder := json.NewDecoder(responseWithoutArray)

	var result = new(AzureTranslationResponse)

	err := decoder.Decode(&result)

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func removeArrayFromResponse(resp *http.Response) io.Reader {

	responseByte, _ := ioutil.ReadAll(resp.Body)

	responseString := string(responseByte)

	responseString = strings.TrimSuffix(responseString, "]")
	responseString = strings.TrimPrefix(responseString, "[")

	output := strings.NewReader(responseString)

	return output
}
