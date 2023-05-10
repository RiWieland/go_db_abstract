package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

type database struct {
	nameSQLiteFile string
	path           string
	instance       *sql.DB
	dbType         string
}

func initializeDb(db database) *sql.DB {
	pathSQLiteFile := db.path + db.nameSQLiteFile
	file, err := os.Create(pathSQLiteFile) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db initialized")
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
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

func (db database) prepareSql(t table) {
	var sqlStatement strings.Builder

	sqlStatement.WriteString("CREATE TABLE IF NOT EXISTS " + t.name + "( ")

	for key, element := range t.columnsType {
		sqlStatement.WriteString(t.columnName[key] + " " + element + ", ")
	}
	fmt.Println(sqlStatement.String())
}

func (db database) reader(tableName []string) []customer {

	var records []customer
	rows, err := db.instance.Query("SELECT * FROM userinfo")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {

		var id int
		var firstName string
		var lastName string
		var age int
		var address string
		var streetAddress string
		var city string
		var state string
		err = rows.Scan(&id, &firstName, &lastName, &age, &address, &streetAddress, &city, &state)
		record := customer{id, firstName + lastName, age, customerAddress{streetAddress, city, state}}
		records = append(records, record)
	}
	rows.Close() //good habit to close

	return records

}
