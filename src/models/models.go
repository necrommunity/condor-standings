package models

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	db *sql.DB
)

//-----------
// Data types
//-----------

// Models for each property represents a database table
type Models struct {
	Users
}

// Close the database
func (*Models) Close() {
	db.Close()
}

//------------------------
// Initialization function
//------------------------

// Init the models
func Init() (*Models, error) {
	// Read the database configuration from environment variables
	// (it was loaded from the .env file in main.go)
	dbHost := os.Getenv("DB_HOST")
	if len(dbHost) == 0 {
		return nil, errors.New("The \"DB_HOST\" environment variable is blank.")
	}
	dbUser := os.Getenv("DB_USER")
	if len(dbUser) == 0 {
		return nil, errors.New("The \"DB_USER\" environment variable is blank.")
	}
	dbPass := os.Getenv("DB_PASS")
	if len(dbPass) == 0 {
		return nil, errors.New("The \"DB_PASS\" environment variable is blank.")
	}
	dbName := os.Getenv("DB_NAME")
	if len(dbPass) == 0 {
		return nil, errors.New("The \"DB_NAME\" environment variable is blank.")
	}

	// Initialize the database
	// (3306 is the default port for MariaDB)
	dsn := dbUser + ":" + dbPass + "@(" + dbHost + ":3306)/" + dbName + "?parseTime=true"
	if v, err := sql.Open("mysql", dsn); err != nil {
		return nil, err
	} else {
		db = v
	}

	// Create the model
	return &Models{}, nil
}
