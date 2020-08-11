package models

type Astro struct {
	Id          int
	ImagePath   string
	Name        string
	Moons       int
	Description string
	Category    string
	Information FisicalInformation
}
