package models

type FisicalInformation struct {
	Mass        string  `json:"mass"`
	Diameter    float64 `json:"diameter"`
	Temperature float64 `json:"temperature"`
	SunDistance float64 `json:"sunDistance"`
}
