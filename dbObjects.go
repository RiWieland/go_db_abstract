package main

import (
	"fmt"
	"strings"
)

// abstract class for db objects:
type dbObject struct {
	name string
}

// Implementing Type
type table struct {
	dbObject
	column
	row
}

// Implementing Type
type view struct {
	dbObject
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
	prepareSQL()
	create()
	// maybe more granular operations: insert, create
}

// Create is used by concrete classes table, view, etc.
func (d *table) create() {
	var sqlStatement strings.Builder

	sqlStatement.WriteString("CREATE TABLE IF NOT EXISTS " + d.name + "( ")

	for key, element := range d.columnsType {
		sqlStatement.WriteString(d.columnsName[key] + " " + element + ", ")
	}
	fmt.Println(sqlStatement.String())
}
