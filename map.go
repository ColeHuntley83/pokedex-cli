package main

import (

	"fmt"
)


func mapCommand(config *config) error {
	fmt.Println("here")
	//if we have not fetched any data yet
	locationsResp, err := config.pokeApiClient.ListLocations(config.nextLocation)
	if err != nil {
		return err
	}
	config.nextLocation = locationsResp.Next
	config.prevLocation = locationsResp.Previous

	for _, val := range locationsResp.Results{
		fmt.Println(val.Name)
	}
	return nil


}