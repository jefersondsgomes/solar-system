package repositories

import (
	"github.com/jefersondsgomes/solar-system-catalog/models"
	"github.com/jefersondsgomes/solar-system-catalog/providers"
	"gorm.io/gorm"
)

const PhysicalData = "PhysicalData"

var db *gorm.DB

func init() {
	providers.Connect()
	db = providers.GetDB()
	db.AutoMigrate(&models.Astro{}, &models.PhysicalData{})
}

func Create(astro models.Astro) (models.Astro, error) {
	var result = db.Create(&astro)
	return astro, result.Error
}

func Get(astro models.Astro) (models.Astro, error) {
	var result = db.Preload(PhysicalData).Find(&astro)
	return astro, result.Error
}

func GetAll() ([]models.Astro, error) {
	var astros []models.Astro
	var result = db.Preload(PhysicalData).Find(&astros)
	return astros, result.Error
}

func Update(astro models.Astro) (models.Astro, error) {
	if err := db.Model(&astro).Updates(&astro).Error; err != nil {
		return astro, err
	}

	err := db.Model(&astro.PhysicalData).Updates(&astro.PhysicalData).Error
	return astro, err
}

func Delete(astro models.Astro) error {
	var result = db.Delete(&astro)
	return result.Error
}
