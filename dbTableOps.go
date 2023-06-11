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
	C []any
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

func ReadEmbbStruct(st interface{}) {
	readEmbbStruct(reflect.ValueOf(st))
}

func readEmbbStruct(val reflect.Value) {

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		// fmt.Println(val.Type().Field(i).Type.Kind())
		f := val.Field(i)
		switch f.Kind() {
		case reflect.Struct:
			readEmbbStruct(f)
		case reflect.Slice:
			for j := 0; j < f.Len(); j++ {
				readEmbbStruct(f.Index(i))
			}
		case reflect.String, reflect.Int:
			fmt.Printf("%v=%v\n", val.Type().Field(i).Name, val.Field(i))

		}
	}
}

func ReadStruct(st interface{}) []reflect.Value {

	var retValues []reflect.Value

	val := reflect.ValueOf(st)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		switch f.Kind() {
		case reflect.Struct:
			ReadStruct(f.Interface())
		case reflect.Slice:
			for j := 0; j < f.Len(); j++ {
				ReadStruct(f.Index(i).Interface())
			}
		default:
			retValues = append(retValues, val)
			fmt.Printf("Call from func: %v=%v\n", val.Type().Field(i).Name, val.Field(i))

		}
	}
	return retValues
}

// only int and string
// func too loopy
func (db database) createTable(t interface{}) {
	val := reflect.ValueOf(t)

	var sqlStatement strings.Builder
	var View bool

	// extract name for Table:
	n := reflect.TypeOf(t).Name()
	if View != true {
		sqlStatement.WriteString("CREATE TABLE IF NOT EXISTS " + n + " (")
	} else {
		sqlStatement.WriteString("CREATE VIEW IF NOT EXISTS " + n + " (")
	}

	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		switch f.Kind() {
		case reflect.Slice:

			val := f.Index(0).Interface()
			s := reflect.ValueOf(val)
			for i := 0; i < s.NumField(); i++ {
				t := s.Field(i)
				switch t.Kind() {
				case reflect.String:
					s := fmt.Sprint(s.Type().Field(i).Name)
					sqlStatement.WriteString(" \"" + s + "\" TEXT,")

				case reflect.Int:
					s := fmt.Sprint(s.Type().Field(i).Name)
					sqlStatement.WriteString(" \"" + s + "\" INTEGER,")
				}
			}
		}
	}

	sz := len(sqlStatement.String())
	ExSql := sqlStatement.String()
	ExSql = ExSql[:sz-1] + ")"

	_, err := db.instance.Exec(ExSql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Executed statement: " + ExSql)

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
