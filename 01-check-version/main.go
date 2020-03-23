package main

import (
	"fmt"
	"github.com/ConquestSolutions/apiv2.examples/conquest_api"
	"os"
)

func GetAPIVersion() (string, error) {
	api, err := conquest_api.New()
	if err != nil {
		return "", err
	}
	response, err := api.SystemService.ApplicationVersion(nil)
	if err != nil {
		return "", err
	}
	return response.GetPayload().ConquestAPIVersion, nil
}

func main() {

	version, err := GetAPIVersion()
	if err != nil {
		println("fatal: " + err.Error())
		os.Exit(1)
	}

	fmt.Printf("You have connected to API %s", version)
}
