package service

import "LMWN-assignment/types"

type CovidCaseService interface {
	GetAllCovidCases() (*types.CovidCasesSummary, error)
}
