package repositories

import (
	"database/sql"

	"../models"
	"../providers"
)

func GetAll() []models.Astro {
	var id, moons int
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

		err2 := db.QueryRow(
			"SELECT Mass, Diameter, Temperature, SunDistance FROM [dbo].[FisicalInformation] WHERE AstroId = @AstroId",
			sql.Named("AstroId", id)).Scan(&mass, &diameter, &temperature, &sunDistance)

		if err2 == sql.ErrNoRows {
			astros = append(astros, models.Astro{id, imagePath, name, moons, description, category, models.FisicalInformation{}})
		} else {
			astros = append(astros, models.Astro{id, imagePath, name, moons, description, category,
				models.FisicalInformation{mass, diameter, temperature, sunDistance}})
		}
	}

	return astros
}

func Get(id int) models.Astro {
	var astro models.Astro
	var moons int
	var imagePath, name, description, category, mass string
	var diameter, temperature, sunDistance float64

	db := providers.SqlConnection()
	defer db.Close()

	err := db.QueryRow(
		"SELECT a.ImagePath, a.Name, a.Moons, a.Description, a.Category, f.Mass, f.Diameter, F.Temperature, F.SunDistance FROM [dbo].[Astro] a LEFT JOIN [dbo].[FisicalInformation] f ON a.Id = f.AstroId WHERE a.Id = @id",
		sql.Named("id", id)).Scan(&imagePath, &name, &moons, &description, &category, &mass, &diameter, &temperature, &sunDistance)

	if err != nil && err == sql.ErrNoRows {
		return astro
	}

	astro = models.Astro{id, imagePath, name, moons, description, category,
		models.FisicalInformation{mass, diameter, temperature, sunDistance}}

	return astro
}

func Create(astro models.Astro) *models.Astro {
	db := providers.SqlConnection()
	defer db.Close()
	return nil

	//insert, err := db.Prepare("INSERT INTO [dbo].[Astro] ()")

}

func Update(id int, astro models.Astro) {
	db := providers.SqlConnection()
	defer db.Close()

}

func Delete(id int) {
	db := providers.SqlConnection()
	defer db.Close()

}
