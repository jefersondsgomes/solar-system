package repositories

import (
	"database/sql"
	"fmt"

	"../models"
	"../providers"
)

func Create(astro models.Astro) error {
	db := providers.SqlConnection()
	defer db.Close()

	var lastInsertId int
	if err := db.QueryRow("INSERT INTO [dbo].[Astro] (Image, Name, Category, Description) OUTPUT Inserted.ID VALUES (@Image, @Name, @Category, @Description)",
		sql.Named("Image", astro.Image),
		sql.Named("Name", astro.Name),
		sql.Named("Category", astro.Category),
		sql.Named("Description", astro.Description)).Scan(&lastInsertId); err != nil {
		return err
	}

	if len(astro.Information.Mass) > 0 &&
		astro.Information.Diameter > 0 {
		insert, err2 := db.Prepare("INSERT INTO [dbo].[FisicalInformation] (AstroId, Mass, Diameter, Temperature, SunDistance) VALUES (@AstroId, @Mass, @Diameter, @Temperature, @SunDistance)")
		if err2 != nil {
			return err2
		}

		_, err3 := insert.Exec(sql.Named("AstroId", lastInsertId), sql.Named("Mass", astro.Information.Mass),
			sql.Named("Diameter", astro.Information.Diameter), sql.Named("Temperature", astro.Information.Temperature),
			sql.Named("SunDistance", astro.Information.SunDistance))
		if err3 != nil {
			return err3
		}
	}

	return nil
}

func GetAll() ([]models.Astro, error) {
	var id int
	var image, name, category, description, mass string
	var diameter, temperature, sunDistance float64

	db := providers.SqlConnection()
	defer db.Close()
	astros := []models.Astro{}

	rows, err := db.Query("SELECT * FROM [dbo].[Astro] ORDER BY Id")
	if err != nil && err == sql.ErrNoRows {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&id, &image, &name, &category, &description)
		if err != nil && err == sql.ErrNoRows {
			return nil, err
		}

		err2 := db.QueryRow(
			"SELECT Mass, Diameter, Temperature, SunDistance FROM [dbo].[FisicalInformation] WHERE AstroId = @AstroId",
			sql.Named("AstroId", id)).Scan(&mass, &diameter, &temperature, &sunDistance)

		if err2 == sql.ErrNoRows {
			astros = append(astros, models.Astro{id, image, name, category, description, models.FisicalInformation{}})
		} else {
			astros = append(astros, models.Astro{id, image, name, category, description,
				models.FisicalInformation{mass, diameter, temperature, sunDistance}})
		}
	}

	return astros, nil
}

func Get(id int) (models.Astro, error) {
	var astro models.Astro
	var image, name, category, description, mass string
	var diameter, temperature, sunDistance float64

	db := providers.SqlConnection()
	defer db.Close()

	err := db.QueryRow(
		"SELECT a.Image, a.Name, a.Category, a.Description, f.Mass, f.Diameter, F.Temperature, F.SunDistance FROM [dbo].[Astro] a LEFT JOIN [dbo].[FisicalInformation] f ON a.Id = f.AstroId WHERE a.Id = @id",
		sql.Named("id", id)).Scan(&image, &name, &category, &description, &mass, &diameter, &temperature, &sunDistance)

	if err != nil && err == sql.ErrNoRows {
		return astro, err
	}

	astro = models.Astro{id, image, name, category, description,
		models.FisicalInformation{mass, diameter, temperature, sunDistance}}

	return astro, nil
}

func Update(id int, astro models.Astro) error {
	db := providers.SqlConnection()
	defer db.Close()

	update1 := fmt.Sprintf("UPDATE [dbo].[Astro] SET Image = '%s', Name = '%s', Category = '%s', Description = '%s' WHERE Id = %d",
		astro.Image, astro.Name, astro.Category, astro.Description, id)

	_, err := db.Exec(update1)
	if err != nil {
		return err
	}

	if (models.FisicalInformation{}) != astro.Information {
		dbAstro, _ := Get(id)

		if (models.FisicalInformation{}) != dbAstro.Information {
			update2 := fmt.Sprintf("UPDATE [dbo].[FisicalInformation] SET Mass = '%s', Diameter = %f, Temperature = %f, SunDistance = %f WHERE AstroId = %d",
				astro.Information.Mass, astro.Information.Diameter, astro.Information.Temperature, astro.Information.SunDistance, id)

			_, err2 := db.Exec(update2)
			if err2 != nil {
				return err2
			}
		} else {
			insert := fmt.Sprintf("INSERT INTO [dbo].[FisicalInformation] (AstroId, Mass, Diameter, Temperature, SunDistance) VALUES (%d, '%s', %f, %f, %f)",
				id, astro.Information.Mass, astro.Information.Diameter, astro.Information.Temperature, astro.Information.SunDistance)

			_, err3 := db.Exec(insert)
			if err3 != nil {
				return err3
			}
		}
	}

	return nil
}

func Delete(id int) error {
	db := providers.SqlConnection()
	defer db.Close()

	delete, err := db.Prepare("DELETE FROM [dbo].[Astro] WHERE Id = @id")
	if err != nil {
		return err
	}

	_, err2 := delete.Exec(sql.Named("id", id))
	if err2 != nil {
		return err2
	}

	return nil
}
