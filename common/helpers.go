package common

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type configuration struct {
	AzureAPIKey string
}

// GetAPIKey gets the API key from a config file
func GetAPIKey() string {
	configuration := configuration{}

	configErr := gonfig.GetConf("TranslatorConfig.json", &configuration)

	if configErr != nil {
		fmt.Println("Error getting config: ", configErr)
	}

	return configuration.AzureAPIKey
}
