package main

import (
	"fmt"

	//"strings"

	"path"

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

var db database

func main() {
	fmt.Println("test")
	var jsonRawStorage rawStorage
	var csvRawStorage rawStorage
	jsonRawStorage.path = "data_json/"
	jsonRawStorage.fileFormat = ".json"

	csvRawStorage.path = "data_csv/customer_20230415.csv"
	csvRawStorage.fileFormat = path.Ext(csvRawStorage.path)

	db.path = "./db/"
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
	customerB.firstname = "Allen"
	customerB.lastname = "Cuck"
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

	customerTableAbstract := table{
		name: "Customer",
		view: false,
	}
	customerTable := customerCollection{
		table: customerTableAbstract,
	}

	orderTableAbstract := table{
		name: "Order",
		view: false,
	}
	orderTable := orderCollection{
		table: orderTableAbstract,
	}
	orderTable.o = []order{ordersA, ordersB}
	customerTable.c = []customer{customerA, customerB}
	fmt.Println(customerTable)

	customerTable.createTable()

}
