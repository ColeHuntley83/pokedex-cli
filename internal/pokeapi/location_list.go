package pokeapi

import (
	"encoding/json"
	"io"
)


func (c *Client) ListLocations(pageUrl *string)(RespLocationList, error){

 	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	resp, err :=  c.httpClient.Get(url)

	if err != nil {
		return RespLocationList{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationList{}, err
	}
	locationList := RespLocationList{}
	err = json.Unmarshal(data, &locationList)
	if err != nil {
		return RespLocationList{}, err
	}
	return locationList, nil
}

