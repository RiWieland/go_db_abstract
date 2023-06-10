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
		}
	}
	return retValues
}

/*
// -> create slice of same type: (is this necessary?)
sliceType := reflect.SliceOf(reflect.TypeOf(k))
slice := reflect.MakeSlice(sliceType, 10, 10)
// -> fill slice with custom elemens:
*/

// only int and string
// on orderCollection -> make it for all tables
func (db database) createTable(t orderCollection) {

	var sqlStatement strings.Builder

	if t.View != true {
		sqlStatement.WriteString("CREATE TABLE IF NOT EXISTS " + t.Name + " (")
	} else {
		sqlStatement.WriteString("CREATE VIEW IF NOT EXISTS " + t.Name + " (")
	}

	for _, order := range t.C {
		//ReadEmbbStruct(orders)
		res := ReadStruct(order)

		for i, j := range res {
			val := j.Field(i)
			switch val.Kind() {
			case reflect.String:

				s := fmt.Sprint(j.Type().Field(i).Name)
				sqlStatement.WriteString(" \"" + s + "\" TEXT,")

			case reflect.Int:
				s := fmt.Sprint(j.Type().Field(i).Name)
				sqlStatement.WriteString(" \"" + s + "\" INTEGER,")
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
