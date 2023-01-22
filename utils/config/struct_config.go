package config

type config struct {
	WongnaiCovidCases string   `yaml:"wongnai_covid_cases"`
	Cors              []string `yaml:"cors"`
	Environment       int8     `yaml:"environment"`
	Port              string   `yaml:"port"`
}
