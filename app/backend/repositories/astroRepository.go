package repositories

import (
	"database/sql"
	"fmt"

	"../models"
	"../providers"
)

func Create(astro models.Astro) {
	db := providers.SqlConnection()
	defer db.Close()

	var lastInsertId int
	if err := db.QueryRow("INSERT INTO [dbo].[Astro] (ImagePath, Name, Moons, Description, Category) OUTPUT Inserted.ID VALUES (@ImagePath, @Name, @Moons, @Description, @Category)",
		sql.Named("ImagePath", astro.ImagePath),
		sql.Named("Name", astro.Name),
		sql.Named("Moons", astro.Moons),
		sql.Named("Description", astro.Description),
		sql.Named("Category", astro.Category)).Scan(&lastInsertId); err != nil {
		panic(err)
	}

	if len(astro.Information.Mass) > 0 &&
		astro.Information.Diameter > 0 {
		insert, err2 := db.Prepare("INSERT INTO [dbo].[FisicalInformation] (AstroId, Mass, Diameter, Temperature, SunDistance) VALUES (@AstroId, @Mass, @Diameter, @Temperature, @SunDistance)")
		if err2 != nil {
			panic(err2)
		}

		_, err3 := insert.Exec(sql.Named("AstroId", lastInsertId), sql.Named("Mass", astro.Information.Mass),
			sql.Named("Diameter", astro.Information.Diameter), sql.Named("Temperature", astro.Information.Temperature),
			sql.Named("SunDistance", astro.Information.SunDistance))
		if err3 != nil {
			panic(err3)
		}
	}
}

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

func Update(id int, astro models.Astro) {
	db := providers.SqlConnection()
	defer db.Close()

	update1 := fmt.Sprintf("UPDATE [dbo].[Astro] SET ImagePath = '%s', Name = '%s', Moons = %d, Description = '%s', Category = '%s' WHERE Id = %d",
		astro.ImagePath, astro.Name, astro.Moons, astro.Description, astro.Category, id)

	_, err := db.Exec(update1)
	if err != nil {
		panic(err)
	}

	if (models.FisicalInformation{}) != astro.Information {
		dbAstro := Get(id)

		if (models.FisicalInformation{}) != dbAstro.Information {
			update2 := fmt.Sprintf("UPDATE [dbo].[FisicalInformation] SET Mass = '%s', Diameter = %f, Temperature = %f, SunDistance = %f WHERE AstroId = %d",
				astro.Information.Mass, astro.Information.Diameter, astro.Information.Temperature, astro.Information.SunDistance, id)

			_, err2 := db.Exec(update2)
			if err2 != nil {
				panic(err2)
			}
		} else {
			insert := fmt.Sprintf("INSERT INTO [dbo].[FisicalInformation] (AstroId, Mass, Diameter, Temperature, SunDistance) VALUES (%d, '%s', %f, %f, %f)",
				id, astro.Information.Mass, astro.Information.Diameter, astro.Information.Temperature, astro.Information.SunDistance)

			_, err3 := db.Exec(insert)
			if err3 != nil {
				panic(err3)
			}
		}
	}
}

func Delete(id int) {
	db := providers.SqlConnection()
	defer db.Close()

	delete, err := db.Prepare("DELETE FROM [dbo].[Astro] WHERE Id = @id")
	if err != nil {
		panic(err)
	}

	delete.Exec(sql.Named("id", id))
}
