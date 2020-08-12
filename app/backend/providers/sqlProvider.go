package providers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var server = "localhost"
var port = 1433
var user = "sa"
var password = "1SSCataloger@"
var database = "SOLAR.SYSTEM.CATALOG"

func SqlConnection() *sql.DB {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	db, err := sql.Open("sqlserver", connectionString)

	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	return db
}
