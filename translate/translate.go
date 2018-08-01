package translate

import (
	httprequests "GoTranslator/httpRequests"
	"fmt"
)

// Translate an input string
func Translate(input string) string {

	translation := httprequests.HTTPPostRequest("https://api.cognitive.microsofttranslator.com/translate?api-version=3.0", input, "fr")

	if translation == nil {
		fmt.Println("Error translating")
		return ""
	}

	//fmt.Println(translation.Translations[0].Text)

	return translation.Translations[0].Text
}
