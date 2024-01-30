package main

import (
	"fmt"
	"log"

	"github.com/PaoloProdossimoLopes/goshop/configs"
)

func main() {
	configurations, loadConfigurationError := configs.LoadConfigurations(".")
	if loadConfigurationError != nil {
		log.Fatalf("Error loading configurations: %v", loadConfigurationError)
		panic(loadConfigurationError)
	}

	fmt.Println(configurations)
}
