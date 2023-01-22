package repository

import (
	"encoding/json"
)

type covidCaseRepositoryMock struct {
	covidCases CovidCase
}

func NewCovidCaseRepositoryMock() covidCaseRepositoryMock {
	// * Mock data
	covidCasesMock := map[string]interface{}{
		"Data": []map[string]interface{}{
			{
				"ConfirmDate":    "2021-06-02",
				"No":             nil,
				"Age":            34,
				"Gender":         nil,
				"GenderEn":       nil,
				"Nation":         nil,
				"NationEn":       "Thailand",
				"Province":       "Ranong",
				"ProvinceId":     50,
				"District":       nil,
				"ProvinceEn":     "Ranong",
				"StatQuarantine": 14,
			},
			{
				"ConfirmDate":    "2021-05-06",
				"No":             nil,
				"Age":            nil,
				"Gender":         "หญิง",
				"GenderEn":       "Female",
				"Nation":         nil,
				"NationEn":       nil,
				"Province":       "Samut Songkhram",
				"ProvinceId":     58,
				"District":       nil,
				"ProvinceEn":     "Samut Songkhram",
				"StatQuarantine": 11,
			},
			{
				"ConfirmDate":    "2021-05-02",
				"No":             nil,
				"Age":            72,
				"Gender":         nil,
				"GenderEn":       nil,
				"Nation":         nil,
				"NationEn":       "China",
				"Province":       "Sisaket",
				"ProvinceId":     62,
				"District":       nil,
				"ProvinceEn":     "Sisaket",
				"StatQuarantine": 20,
			},
			{
				"ConfirmDate":    nil,
				"No":             nil,
				"Age":            27,
				"Gender":         "หญิง",
				"GenderEn":       "Female",
				"Nation":         nil,
				"NationEn":       "Thailand",
				"Province":       nil,
				"ProvinceId":     nil,
				"District":       nil,
				"ProvinceEn":     nil,
				"StatQuarantine": 6,
			},
		},
	}

	covidCasesData := new(CovidCase)

	// * Marshall to normal json
	jsonbody, _ := json.Marshal(covidCasesMock)

	// * Unmarshall to struct
	_ = json.Unmarshal(jsonbody, covidCasesData)

	return covidCaseRepositoryMock{covidCases: *covidCasesData}
}

func (r covidCaseRepositoryMock) GetCovidCases() (*CovidCase, error) {
	return &r.covidCases, nil
}
