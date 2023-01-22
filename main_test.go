package main

import (
	"LMWN-assignment/repository"
	"LMWN-assignment/router"
	"LMWN-assignment/service"
	"LMWN-assignment/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestCovidSummaryRouteAvailableHandler(t *testing.T) {
	gonic := gin.Default()

	// * Init router
	endpoints := gonic.Group("/")
	router.Router(endpoints)

	request, _ := http.NewRequest("GET", "/covid/summary", nil)
	writer := httptest.NewRecorder()
	gonic.ServeHTTP(writer, request)
	assert.Equal(t, 200, writer.Code)
}

func TestCovidSummaryMock(t *testing.T) {
	expectedResult := types.CovidCasesSummary{
		Province: map[string]int64{
			"Chai Nat":        1,
			"Ranong":          2,
			"Samut Songkhram": 1,
			"Sisaket":         1,
		},
		AgeGroup: map[string]int64{
			"0-30":  2,
			"31-60": 1,
			"61+":   2,
			"N/A":   1,
		},
	}

	// * Register service
	covidCaseRepositoryMock := repository.NewCovidCaseRepositoryMock()
	covidCaseService := service.NewCovidCaseService(covidCaseRepositoryMock)

	// * Mocked cases
	caseMocked, _ := covidCaseService.GetAllCovidCases()

	// * Comparing result
	assert.Equal(t, expectedResult, caseMocked)
}

func TestCovidNumberOfProvince(t *testing.T) {
	// * Register service
	covidCaseRepositoryMock := repository.NewCovidCaseRepositoryMock()
	covidCaseService := service.NewCovidCaseService(covidCaseRepositoryMock)

	// * Mocked cases
	caseMocked, _ := covidCaseService.GetAllCovidCases()

	// * Expected of age group
	expectedProvinceGroupLength := 4

	// * Length of province group
	var caseMockedProvinceLength = 0
	for _ = range caseMocked.Province {
		caseMockedProvinceLength++
	}

	// * Comparing result
	assert.Equal(t, expectedProvinceGroupLength, caseMockedProvinceLength)
}

func TestCovidNumberOfCases(t *testing.T) {
	// * Register service
	covidCaseRepositoryMock := repository.NewCovidCaseRepositoryMock()
	covidCaseService := service.NewCovidCaseService(covidCaseRepositoryMock)

	// * Mocked cases
	caseMocked, _ := covidCaseService.GetAllCovidCases()

	// * Expected of age group
	expectedAgeGroupLength := 6

	// * Length of age group
	var caseMockedAgeGroupLength = 0
	for k := range caseMocked.AgeGroup {
		caseMockedAgeGroupLength += int(caseMocked.AgeGroup[k])
	}

	// * Comparing result
	assert.Equal(t, expectedAgeGroupLength, caseMockedAgeGroupLength)
}
