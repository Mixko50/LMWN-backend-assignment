package router

import (
	"LMWN-assignment/handler"
	"LMWN-assignment/repository"
	"LMWN-assignment/service"
	"LMWN-assignment/utils/custom_http"
	"github.com/gin-gonic/gin"
)

func Router(router gin.IRouter) {

	// * Register service
	covidCaseRepository := repository.NewCovidCaseRepositoryAPI(custom_http.HttpClient)

	covidCaseRepositoryMock := repository.NewCovidCaseRepositoryMock()
	_ = covidCaseRepositoryMock

	covidCaseService := service.NewCovidCaseService(covidCaseRepository)
	covidCaseHandler := handler.NewCovidCaseHandler(covidCaseService)

	// * Routes
	router.GET("/covid/summary", covidCaseHandler.GetAllCovidCases)
}
