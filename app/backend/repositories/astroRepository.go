package repositories

import (
	"database/sql"

	"../models"
	"../providers"
)

func GetAll() []models.Astro {
	var id, moons, informationId, astroFk int
	var imagePath, name, description, category, mass string
	var diameter, temperature, sunDistance float64

	db := providers.SqlConnection()
	defer db.Close()
	astros := []models.Astro{}

	rows, err := db.Query("SELECT * FROM [dbo].[Astro] ORDER BY Id")
	if err != nil && err == sql.ErrNoRows {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&id, &imagePath, &name, &moons, &description, &category)
		if err != nil && err == sql.ErrNoRows {
			return nil
		}

		err2 := db.QueryRow("SELECT * FROM [dbo].[FisicalInformation] WHERE AstroId = @AstroId", sql.Named("AstroId", id)).Scan(&informationId, &astroFk, &mass, &diameter, &temperature, &sunDistance)
		if err2 == sql.ErrNoRows {
			astros = append(astros, models.Astro{id, imagePath, name, moons, description, category, models.FisicalInformation{}})
		} else {
			information := models.FisicalInformation{mass, diameter, temperature, sunDistance}
			astros = append(astros, models.Astro{id, imagePath, name, moons, description, category, information})
		}
	}

	return astros
}

func Get(id int) models.Astro {
	var astro models.Astro
	var information models.FisicalInformation
	var moons, informationId, astroFk int
	var imagePath, name, description, category, mass string
	var diameter, temperature, sunDistance float64

	db := providers.SqlConnection()
	defer db.Close()

	err := db.QueryRow("SELECT TOP 1 * FROM [dbo].[Astro] WHERE Id = @id", sql.Named("id", id)).Scan(&id, &imagePath, &name, &moons, &description, &category)
	if err != nil && err == sql.ErrNoRows {
		return astro
	}

	err2 := db.QueryRow("SELECT TOP 1 * FROM [dbo].[FisicalInformation] WHERE AstroId = @id", sql.Named("id", id)).Scan(&informationId, &astroFk, &mass, &diameter, &temperature, &sunDistance)
	if err2 != sql.ErrNoRows {
		information = models.FisicalInformation{mass, diameter, temperature, sunDistance}
	}

	astro = models.Astro{id, imagePath, name, moons, description, category, information}
	return astro
}

func Create(astro models.Astro) {
	db := providers.SqlConnection()
	defer db.Close()

}

func Update(id int, astro models.Astro) {
	db := providers.SqlConnection()
	defer db.Close()

}

func Delete(id int) {
	db := providers.SqlConnection()
	defer db.Close()

}
