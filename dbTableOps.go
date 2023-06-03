package main

import (
	"fmt"
	"strings"
)

// Implementing Type
type table struct {
	name string
	view bool
	column
	row
}

// Column defines the datataypes
type column struct {
	name        string
	columnsType []string
	columnsName []string
}

// rows contain the values
type row struct {
	id        int
	rowValues []string
}

// Abstract Interface
type operation interface {
	filterTable()
	createTable()
	// Is Interface neccessary?
}

// Create returns SQL Statement
// implement sql-types via mapping table
// how to connect the datatypes of go objects?
func (t table) createTable() {

	var sqlStatement strings.Builder

	if t.view != true {
		sqlStatement.WriteString("CREATE TABLE IF NOT EXISTS " + t.name + "( ")
	} else {
		sqlStatement.WriteString("CREATE VIEW IF NOT EXISTS " + t.name + "( ")
	}

	for key, element := range t.columnsType {
		sqlStatement.WriteString(t.columnsName[key] + " " + element + ", ")
		fmt.Println(t.columnsName[key] + " " + element + ", ")
	}
	fmt.Println(sqlStatement.String())
	_, err := db.instance.Exec(sqlStatement.String())
	if err != nil {
		fmt.Println(err)
	}

}

/*
// function must query table and convert it to the go object with mapped types
// how to integrate the object
func (t table) query() table {

	var records []customer
	rows, err := db.instance.Query("SELECT * FROM " + t.name)
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
*/

func (t table) insert() {
	var sqlStatement strings.Builder

	sqlStatement.WriteString("INSERT INTO " + t.name + "( ")

	for key, element := range t.columnsType {
		sqlStatement.WriteString(t.columnsName[key] + " " + element + ", ")
	}
	_, err := db.instance.Exec(sqlStatement.String())
	if err != nil {
		fmt.Println(err)
	}
}

func (t table) filter() {

}

// function joins other table on specified column
func (t table) leftJoin(tableJoin table, col column) {

}
