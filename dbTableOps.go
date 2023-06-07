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

// only int and string
func (db database) createTable(t customerCollection) {

	var sqlStatement strings.Builder

	if t.Table.View != true {
		sqlStatement.WriteString("CREATE TABLE IF NOT EXISTS " + t.Name + "( ")
	} else {
		sqlStatement.WriteString("CREATE VIEW IF NOT EXISTS " + t.Name + "( ")
	}

	coll := t.C
	f := coll[0]
	e := reflect.ValueOf(&f).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varValue := e.Field(i).Interface()
		//varType := e.Type().Field(i).Type

		switch varValue.(type) {
		case int:
			i := fmt.Sprint(varName)
			sqlStatement.WriteString("\"" + i + "\"  INTEGER, ")

		case string:
			i := fmt.Sprint(varName)
			sqlStatement.WriteString("\"" + i + "\"  TEXT, ")
		}
	}

	_, err := db.instance.Exec(sqlStatement.String())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Executed statement: " + sqlStatement.String())

}

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
