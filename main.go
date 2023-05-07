package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"

	//"strings"

	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

/*
To-Do:
- add type mapping: then maps for tables does not need anymore "string"
- for function "prepareSQL":
- - should work for different types (view, table) -> interface?
- - need type of database for different SQL styles
- - implement type mapping from Go-Types to DB-types
- table-struct: how to best create table columns and types?

*/
/*
how to structure classes for different db's
- Database (abstract)
- - Postgres (or abstract on this level)
- - SQLite
- - ...

- methods for DB: (interfaces)
- - execute
- - prepare
- - read/ query


Classes
Tables: (tables abstract?):
- LandingTables
- tables
- views

methods:
- insert
- create
- merge
- update

Datatypes
- Postgres
	- VARCHAR
	- STRING
	- ...

// Abstract Class:
// Tables:
// -> uses method "write to db"

// Sub Classes:
// special tables;
// implement interfaces for specific types



*/
type handler interface {
	//logger()
	writer()
	reader()
}

type rawStorage struct {
	path       string
	fileFormat string
	totalSize  float64
}

type customerAddress struct {
	street string
	city   string
	state  string
}

type customer struct {
	id      int
	name    string
	age     int
	address customerAddress
}

func main() {

	var jsonRawStorage rawStorage
	var csvRawStorage rawStorage
	jsonRawStorage.path = "data_json/"
	jsonRawStorage.fileFormat = ".json"

	csvRawStorage.path = "data_csv/customer_20230415.csv"
	csvRawStorage.fileFormat = path.Ext(csvRawStorage.path)

	var db database
	db.path = "db/"
	db.nameSQLiteFile = "sqlite-database.db"
	db.instance = initializeDb(db)
	defer db.instance.Close() // Defer Closing the database

	// table:
	var customer table
	customer.name = "Customer"

	customer.columnName = []string{"id", "firstname", "lastName", "age", "address", "streetAddress", "city", "state"}
	customer.columnsType = []string{"integer", "TEXT", "TEXT", "TEXT", "TEXT", "TEXT", "TEXT", "TEXT"}

	db.prepareSql(customer)
	fileExtension := jsonRawStorage.fileFormat
	if fileExtension == ".csv" {
		//extract := readCsvFile(filePath)
		extract_json := csvRawStorage.reader("test")
		fmt.Println(extract_json)
	} else if fileExtension == ".json" {
		extract_json := jsonRawStorage.reader("/customer_20230412.json")
		fmt.Println(extract_json)
	}
}

func readCsvFile(filePath string) []customer {
	recordsCustomer := []customer{}
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data_extract, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	fmt.Println(data_extract)
	_, data_content := data_extract[0], data_extract[1:]
	for _, data := range data_content {
		fmt.Println(data[0])
		age_int, err := strconv.Atoi(data[0])
		if err != nil {
			log.Fatal("can not convert to int: ", age_int, err)
		}
		recordsCustomer = append(recordsCustomer, customer{age_int, data[1], 7, customerAddress{"", "", ""}})
		fmt.Println(recordsCustomer)
	}
	return recordsCustomer
}

func (r rawStorage) reader(fileNames string) customer {
	filePath := r.path + fileNames
	fmt.Println(fileNames)
	customerJson, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer customerJson.Close()

	byteJson, _ := ioutil.ReadAll(customerJson)

	var records map[string]interface{} // -> if we dont know the structure of the json
	err = json.Unmarshal(byteJson, &records)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	//build_str :=
	for key, element := range records {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	nestedMap := records["address"].(map[string]interface{})
	customerReturn := customer{1, records["lastName"].(string), int(records["age"].(float64)), customerAddress{nestedMap["streetAddress"].(string), nestedMap["city"].(string), nestedMap["state"].(string)}}

	return customerReturn
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

func (r rawStorage) writer(filename string, dataCustomer string) {

	file, _ := json.MarshalIndent(dataCustomer, "", " ")
	filePath := r.path + filename
	_ = ioutil.WriteFile(filePath, file, 0644)

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

type database struct {
	nameSQLiteFile string
	path           string
	instance       *sql.DB
	dbType         string
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

// Every different databse systems needs individual function for creating Statement
// -> different sql languages

// special tables need to inherit execute from the database

// Create Statement for RawTables
type dbOperation interface {
	// usage of interface: for raw the Datatype needs conversion, for tables it does not
	createTable()
	insertTable()
}

// pointer for db?

/*
	createCustomerTableSQL := `CREATE TABLE IF NOT EXISTS CUSTOMER (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"firstName" TEXT,
		"lastName" TEXT,
		"age" TEXT
		"address" TEXT
		"streetAddress" TEXT,
        "city" TEXT,
        "state" "Louisiana" TEXT
	  );` // SQL Statement for Create Table
*/
