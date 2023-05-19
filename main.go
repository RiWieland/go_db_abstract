package main

import (
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
// Every different databse systems needs individual function for creating Statement
// -> different sql languages

// special tables need to inherit execute from the database

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
	db.instance = db.initializeDb()
	defer db.instance.Close() // Defer Closing the database

	// table:
	var customerA customer
	var customerB customer

	var ordersA order
	var ordersB order

	customerA.id = 1
	customerA.firstname = "Jose"
	customerA.lastname = "Al"
	customerA.age = 36

	customerB.id = 2
	customerB.firstname = "Jose"
	customerB.lastname = "Al"
	customerB.age = 36

	ordersA.id = 1
	ordersA.firstname = "Jose"
	ordersA.lastname = "Al"
	ordersA.amount = 5
	ordersA.shipped = true

	ordersB.id = 2
	ordersB.firstname = "Bold"
	ordersB.lastname = "Eric"
	ordersB.amount = 1
	ordersB.shipped = true

	var orderTable orderCollection
	orderTable.name = "Orders"
	orderTable.o = []order{ordersA, ordersB}

	var customerTable customerCollection
	customerTable.name = "Customner"
	customerTable.c = []customer{customerA, customerB}

	customerTable.create()
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
