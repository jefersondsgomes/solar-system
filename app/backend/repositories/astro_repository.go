package repositories

import (
	"github.com/jefersondsgomes/universe-catalog/entities"
	"github.com/jefersondsgomes/universe-catalog/providers"
	"github.com/jefersondsgomes/universe-catalog/utils"
	"gorm.io/gorm"
)

const PhysicalData = "PhysicalData"

var db *gorm.DB

func init() {
	providers.Connect()
	db = providers.GetDB()
	db.AutoMigrate(&entities.Astro{}, &entities.PhysicalData{})
}

func Create(astro entities.Astro) (entities.Astro, error) {
	err := db.Create(&astro).Error
	return astro, err
}

func Get(astro entities.Astro) (entities.Astro, error) {
	err := db.Preload(PhysicalData).Find(&astro).Error
	return astro, err
}

func GetAll(pagination utils.Pagination) ([]entities.Astro, error) {
	var astros []entities.Astro
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Limit(pagination.Limit).Offset(offset)
	err := queryBuider.Preload(PhysicalData).Find(&astros).Error
	return astros, err
}

func Update(astro entities.Astro) (entities.Astro, error) {
	err := db.Save(&astro).Error
	return astro, err
}

func Delete(astro entities.Astro) error {
	err := db.Delete(&astro).Error
	return err
}
