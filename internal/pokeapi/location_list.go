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

	// 1. make the GET request using c.httpClient
	res, err := c.httpClient.Get(url)
	if err != nil {
		return Response{}, fmt.Errorf("Error: %v", err)
	}

	// 2. defer close the body
	defer res.Body.Close()
	// 3. io.ReadAll the body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{},fmt.Errorf("Error: %v", err)
	}

	// 4. unmarshal into a Response
	var response Response
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return Response{},fmt.Errorf("Error: %v", err)
	}


	return response, nil
}
