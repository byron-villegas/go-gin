package controller

import (
	"go-gin/service"

	"github.com/gin-gonic/gin"
)

type StarWarsController struct{}

var starWarsService = service.NewStarWarsService()

// GetPeople godoc
// @Summary Get all people
// @Description Get all people from Star Wars API
// @Tags StarWars
// @Accept json
// @Produce json
// @Success 200 {object} dto.PeopleDto
// @Router /starwars/people [get]
func (swc *StarWarsController) GetPeople(c *gin.Context) {
	peoples, err := starWarsService.FindAllPeople()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, peoples)
}
