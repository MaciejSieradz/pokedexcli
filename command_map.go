package main

import (
	"fmt"
)

func commandMapf(config *config) error {
	locationResp, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationResp.Next
	config.previousLocationsURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(config *config) error {
	if config.previousLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationResp, err := config.pokeapiClient.ListLocations(config.previousLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationResp.Next
	config.previousLocationsURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}
