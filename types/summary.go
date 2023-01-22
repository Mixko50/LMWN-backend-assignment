package types

import "math"

type CovidCasesSummary struct {
	Province map[string]int64 `json:"Province"`
	AgeGroup map[string]int64 `json:"AgeGroup"`
}

type AgeGroupConditionStruct struct {
	Title string
	Min   int64
	Max   int64
	Count int64
}

var AgeGroupCondition = []AgeGroupConditionStruct{
	{
		Title: "0-30",
		Min:   0,
		Max:   30,
		Count: 0,
	},
	{
		Title: "31-60",
		Min:   31,
		Max:   60,
		Count: 0,
	},
	{
		Title: "61+",
		Min:   61,
		Max:   math.MaxInt64,
		Count: 0,
	},
}
