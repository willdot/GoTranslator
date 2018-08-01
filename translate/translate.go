package translate

import (
	httprequests "GoTranslator/httpRequests"
	"fmt"
)

// Translate an input string by calling the Azure service
func Translate(input, language string) {

	translations := httprequests.HTTPPostRequest("https://api.cognitive.microsofttranslator.com/translate?api-version=3.0", input, language)

	if translations == nil {
		fmt.Println("Error translating")
		return
	}

	for _, translation := range translations.Translations {
		fmt.Println(translation.To + ": " + translation.Text)
	}
}
