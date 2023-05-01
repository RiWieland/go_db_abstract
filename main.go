package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library

	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

type handler interface {
	//logger()
	writer()
	reader()
}

type database struct {
	nameSQLiteFile string
	path           string
	instance       *sql.DB
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
	createTable(db.instance)  // Create Database Tables

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

func createTable(db *sql.DB) {
	createCustomerTableSQL := `CREATE TABLE IF NOT EXISTS CUSTOMER (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"firstName" TEXT,
		"lastName" TEXT,
		"name" TEXT,
		"age" TEXT
		"address" TEXT		
		"streetAddress" TEXT,
        "city" TEXT,
        "state" "Louisiana" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Create CUSTOMER table...")
	statement, err := db.Prepare(createCustomerTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("CUSTOMER table created")
}
