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

func Create(astro models.Astro) models.Astro {
	db.Create(&astro)
	return astro
}

func Get(id string) models.Astro {
	var astro models.Astro
	db.Where("Id = ?", id).Find(&astro)
	return astro
}

func GetAll() []models.Astro {
	var astros []models.Astro
	db.Find(&astros)
	return astros
}

func Update(id string, astro models.Astro) models.Astro {
	var dbAstro models.Astro
	db.Where("Id = ?", id).Find(&dbAstro)

	if dbAstro.ID == 0 {
		return dbAstro
	}

	db.Model(&astro).Where("Id = ?", id).Updates(map[string]interface{}{"Name": astro.Name, "Category": astro.Category, "Description": astro.Description, "Image": astro.Image, "Mass": astro.Data.Mass, "Temperature": astro.Data.Temperature, "Gravity": astro.Data.Gravity, "SunDistance": astro.Data.SunDistance, "OrbitalPeriod": astro.Data.OrbitalPeriod})
	astro.ID = dbAstro.ID
	return astro
}

func Delete(id string) {
	var astro models.Astro
	db.Where("Id = ?", id).Delete(astro)
}
