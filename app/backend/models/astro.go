package models

type Astro struct {
	Id          int                `json:"id"`
	Image       string             `json:"image"`
	Name        string             `json:"name"`
	Category    string             `json:"category"`
	Description string             `json:"description"`
	Information FisicalInformation `json:"information"`
}
