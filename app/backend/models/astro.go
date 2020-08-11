package models

type Astro struct {
	Id          int                `json:"id"`
	ImagePath   string             `json:"imagePath"`
	Name        string             `json:"name"`
	Moons       int                `json:"moons"`
	Description string             `json:"description"`
	Category    string             `json:"category"`
	Information FisicalInformation `json:"information"`
}
