package models

import "time"

type Astro struct {
	ID          uint64    `json:"id" gorm:"primaryKey;auto_increment"`
	Name        string    `json:"name" binding:"required" gorm:"type:varchar(50);not null"`
	Category    string    `json:"category" binding:"required" gorm:"type:varchar(50)"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	Image       string    `json:"image" gorm:"type:varchar(255)"`
	Data        AstroData `json:"data" gorm:"embedded"`
	Inserted    time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}
