package starwarsapi

import (
	"encoding/json"
	"errors"
	"go-gin/client/starwarsapi/dto"
	"net/http"
	"time"
)

// Cliente HTTP para consumir el servicio externo
type StarWarsApiClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewStarWarsApiClient() *StarWarsApiClient {
	return &StarWarsApiClient{
		BaseURL:    "https://swapi.info/api",
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *StarWarsApiClient) FindAllPeople() ([]*dto.PeopleDto, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/people", nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("product not found")
	}

	var peoples []dto.PeopleDto

	if err := json.NewDecoder(resp.Body).Decode(&peoples); err != nil {
		return nil, err
	}

	peoplePtrs := make([]*dto.PeopleDto, len(peoples))
	for i := range peoples {
		peoplePtrs[i] = &peoples[i]
	}

	return peoplePtrs, nil
}
