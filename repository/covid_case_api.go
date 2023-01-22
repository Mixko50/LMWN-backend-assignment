package repository

import (
	"LMWN-assignment/utils/config"
	"encoding/json"
	"io"
	"net/http"
)

type covidCasesRepositoryAPI struct {
	http *http.Client
}

func NewCovidCaseRepositoryAPI(http *http.Client) covidCasesRepositoryAPI {
	return covidCasesRepositoryAPI{http: http}
}

func (r covidCasesRepositoryAPI) GetCovidCases() (*CovidCase, error) {
	// * Send request
	req, err := r.http.Get(config.C.WongnaiCovidCases)
	if err != nil {
		return nil, err
	}

	// * Read response body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var covidCases = new(CovidCase)

	// * Unmarshall data
	err = json.Unmarshal(body, covidCases)
	if err != nil {
		return nil, err
	}

	return covidCases, nil
}
