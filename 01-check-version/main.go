package main

import (
	"fmt"
	"os"

	"github.com/ConquestSolutions/conquest_api.examples/conquest_api"
)

func GetAPIVersion() (string, error) {
	api, err := conquest_api.New()
	if err != nil {
		return "", err
	}
	response, err := api.SystemService.ApplicationVersion(nil, nil)
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
