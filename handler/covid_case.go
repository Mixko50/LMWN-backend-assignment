package handler

import (
	"LMWN-assignment/service"
	"github.com/gin-gonic/gin"
)

type covidCaseHandler struct {
	covidCaseService service.CovidCaseService
}

func NewCovidCaseHandler(covidCaseService service.CovidCaseService) covidCaseHandler {
	return covidCaseHandler{covidCaseService: covidCaseService}
}

func (h covidCaseHandler) GetAllCovidCases(c *gin.Context) {
	covidCases, err := h.covidCaseService.GetAllCovidCases()
	if err != nil {
		c.JSON(500, map[string]bool{
			"success": false,
		})
	}

	c.JSON(200, covidCases)
}
