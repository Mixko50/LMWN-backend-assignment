package mappers

import (
	"LMWN-assignment/repository"
	"LMWN-assignment/types"
)

func ProvinceCount(cases []*repository.CovidCaseDetail) map[string]int64 {
	var provinces = make(map[string]int64)
	for _, covidCase := range cases {
		if covidCase.ProvinceEn != nil {
			provinces[*covidCase.ProvinceEn]++
		}
	}
	return provinces
}

func AgeGroupCount(cases []*repository.CovidCaseDetail) map[string]int64 {
	var ageGroup = make(map[string]int64)
	for _, covidCase := range cases {
		for _, condition := range types.AgeGroupCondition {
			if covidCase.Age != nil && *covidCase.Age >= condition.Min && *covidCase.Age <= condition.Max {
				ageGroup[condition.Title]++
			} else if covidCase.Age == nil {
				ageGroup["N/A"]++
				break
			}
		}
	}
	return ageGroup
}
