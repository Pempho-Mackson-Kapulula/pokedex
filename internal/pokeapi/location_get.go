package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if cachedData, ok := c.cache.Get(url); ok{
		var response Location
		if err := json.Unmarshal(cachedData, &response); err != nil {
			return Location{},fmt.Errorf("Error: %v", err)
		}

		return response, nil
	}

	// make the GET request using c.httpClient
	res, err := c.httpClient.Get(url)
	if err != nil {
		return Location{}, fmt.Errorf("Error: %v", err)
	}
	defer res.Body.Close()
	
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{},fmt.Errorf("Error: %v", err)
	}

	c.cache.Add(url, bodyBytes)

	// 4. unmarshal into a location
	var location Location
	if err := json.Unmarshal(bodyBytes, &location); err != nil {
		return Location{},fmt.Errorf("Error: %v", err)
	}

	return location, nil
}
