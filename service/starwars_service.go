package service

import (
	client "go-gin/client/starwarsapi"
	"go-gin/client/starwarsapi/dto"
)

type StarWarsService struct {
	Client *client.StarWarsApiClient
}

func NewStarWarsService() *StarWarsService {
	return &StarWarsService{
		Client: client.NewStarWarsApiClient(),
	}
}

func (s *StarWarsService) FindAllPeople() ([]*dto.PeopleDto, error) {
	return s.Client.FindAllPeople()
}
