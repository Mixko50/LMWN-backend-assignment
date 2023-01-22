package service

import (
	"LMWN-assignment/logs"
	"LMWN-assignment/mappers"
	"LMWN-assignment/repository"
	"LMWN-assignment/types"
)

type covidCaseService struct {
	covidCaseRepository repository.CovidCaseRepository
}

func NewCovidCaseService(covidCaseRepository repository.CovidCaseRepository) covidCaseService {
	return covidCaseService{covidCaseRepository: covidCaseRepository}
}

func (s covidCaseService) GetAllCovidCases() (*types.CovidCasesSummary, error) {
	covidCases, err := s.covidCaseRepository.GetCovidCases()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// * Define the type
	var covidSummary = new(types.CovidCasesSummary)

	// * Count number of cases in each province
	var provinceGroup = mappers.ProvinceCount(covidCases.Data)
	covidSummary.Province = provinceGroup

	// * Count number of cases in each age group
	var ageGroup = mappers.AgeGroupCount(covidCases.Data)
	covidSummary.AgeGroup = ageGroup

	return covidSummary, nil
}
