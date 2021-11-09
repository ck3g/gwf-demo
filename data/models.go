package data

import (
	"database/sql"
	"os"

	upperDB "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper upperDB.Session

type Models struct {
	// any models inserted here and in the New function
	// are easily accessible throughout the entire application
}

func New(databasePool *sql.DB) Models {
	db = databasePool

	if os.Getenv("DATABASE_TYPE") == "mysql" || os.Getenv("DATABASE_TYPE") == "mariadb" {
		upper, _ = mysql.New(databasePool)
	} else {
		// Postgres
		upper, _ = postgresql.New(databasePool)
	}

	return Models{}
}
