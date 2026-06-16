package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (Response, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedData, ok := c.cache.Get(url); ok{
		var response Response
		if err := json.Unmarshal(cachedData, &response); err != nil {
			return Response{},fmt.Errorf("Error: %v", err)
		}

		return response, nil
	}
	// make the GET request using c.httpClient
	res, err := c.httpClient.Get(url)
	if err != nil {
		return Response{}, fmt.Errorf("Error: %v", err)
	}
	defer res.Body.Close()
	
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{},fmt.Errorf("Error: %v", err)
	}

	c.cache.Add(url, bodyBytes)

	// 4. unmarshal into a Response
	var response Response
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return Response{},fmt.Errorf("Error: %v", err)
	}

	return response, nil
}
