package translate

import httprequests "GoTranslator/httpRequests"

// Translate an input string
func Translate(input string) string {

	//conf := common.GetAPIKey()

	httprequests.HTTPPostRequest("https://api.cognitive.microsofttranslator.com/translate?api-version=3.0", "bonjour", "")
	//common.HTTPGetRequest("http://localhost:8080/api/todos/1")

	return "Some translation here"
}
