package repositories

import (
	"solar-system/models"
	"solar-system/providers"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	providers.Connect()
	db = providers.GetDB()
	db.AutoMigrate(&models.Astro{})
}

func Create(astro models.Astro) (models.Astro, error) {
	var result = db.Create(&astro)
	return astro, result.Error
}

func Get(id string) (models.Astro, error) {
	var astro models.Astro
	var result = db.Where("Id = ?", id).Find(&astro)
	return astro, result.Error
}

func GetAll() ([]models.Astro, error) {
	var astros []models.Astro
	var result = db.Find(&astros)
	return astros, result.Error
}

func Update(id string, astro models.Astro) (models.Astro, error) {
	var dbAstro models.Astro
	db.Where("Id = ?", id).Find(&dbAstro)

	if dbAstro.ID == 0 {
		return dbAstro, nil
	}

	var result = db.Model(&astro).Where("Id = ?", id).Updates(map[string]interface{}{"Name": astro.Name, "Category": astro.Category, "Description": astro.Description, "Image": astro.Image, "Mass": astro.Data.Mass, "Temperature": astro.Data.Temperature, "Gravity": astro.Data.Gravity, "SunDistance": astro.Data.SunDistance, "OrbitalPeriod": astro.Data.OrbitalPeriod})
	astro.ID = dbAstro.ID
	return astro, result.Error
}

func Delete(id string) error {
	var astro models.Astro
	var result = db.Where("Id = ?", id).Delete(astro)
	return result.Error
}
