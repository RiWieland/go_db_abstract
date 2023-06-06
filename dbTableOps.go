package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Implementing Type
type Table struct {
	Name string
	View bool
	Column
	Row
}

// Column defines the datataypes
type Column struct {
	name        string
	columnsType []string
	columnsName []string
}

// rows contain the values
type Row struct {
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
// how to reflect on concrete type /not abstract? -> implement a method on table to
// export for concrete class?
func (db database) createTable(t customerCollection) {

	var sqlStatement strings.Builder
	//t := i.(customerCollection)

	if t.Table.View != true {
		sqlStatement.WriteString("CREATE TABLE IF NOT EXISTS " + t.Name + "( ")
	} else {
		sqlStatement.WriteString("CREATE VIEW IF NOT EXISTS " + t.Name + "( ")
	}

	// 1. extract collection customers
	// 2. infer type of fields in customers

	coll := t.C
	f := coll[0]
	//ty := reflect.TypeOf(f)
	e := reflect.ValueOf(&f).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
	}

	fmt.Println(sqlStatement.String())
	_, err := db.instance.Exec(sqlStatement.String())
	if err != nil {
		fmt.Println(err)
	}

}

func (t Table) reflectStruct() reflect.Value {
	var f reflect.Value
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	return f

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

func (t Table) insert() {
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

func (t Table) filter() {

}

// function joins other table on specified column
func (t Table) leftJoin(tableJoin Table, col Column) {

}
