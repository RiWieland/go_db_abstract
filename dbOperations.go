package main

import (
	"database/sql"
	"log"
	"os"
)

type database struct {
	nameSQLiteFile string
	path           string
	instance       *sql.DB
	dbType         string
}

func (db database) initializeDb() *sql.DB {
	pathSQLiteFile := db.path + db.nameSQLiteFile
	file, err := os.Create(pathSQLiteFile) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db initialized")
	sqliteDatabase, _ := sql.Open("sqlite3", pathSQLiteFile) // Open the created SQLite File
	if err != nil {
		log.Fatal(err.Error())
	}

	return sqliteDatabase
}

// does every database needs own execute?
func (db database) execute(sqlStatement string) {

	log.Println("pereparing db access...")
	statement, err := db.instance.Prepare(sqlStatement) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Statement executed")
}
