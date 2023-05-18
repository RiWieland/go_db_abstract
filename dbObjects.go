package main

import (
	"fmt"
	"strings"
)

// abstract class for db objects:

// Abstract Concrete Type
type dbObject struct {
	name        string
	columnName  []string
	columnsType []string
}

type table struct {
	dbObject
	columnValue []string
}

type view struct {
	dbObject
	columnValue []string
}

// Abstract Interface
type operation interface {
	prepareSQL()
	create()
	// maybe more granular operations: insert, create
}

// Create is used by concrete classes table, view, etc.
func (d *dbObject) create() {
	var sqlStatement strings.Builder

	sqlStatement.WriteString("CREATE TABLE IF NOT EXISTS " + d.name + "( ")

	for key, element := range d.columnsType {
		sqlStatement.WriteString(d.columnName[key] + " " + element + ", ")
	}
	fmt.Println(sqlStatement.String())
}
