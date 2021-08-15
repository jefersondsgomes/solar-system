package models

type AstroData struct {
	Mass          float64 `json:"mass" gorm:"default:0;not null"`
	Temperature   float64 `json:"temperature" gorm:"default:0;not null"`
	Gravity       float64 `json:"gravity" gorm:"default:0;not null"`
	SunDistance   float64 `json:"sunDistance" gorm:"default:0;not null"`
	OrbitalPeriod int     `json:"orbitalPeriod" gorm:"default:0;not null"`
}
