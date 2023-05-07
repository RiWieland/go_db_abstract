package main

import (
	"fmt"
	"strings"
)

func (db database) prepareSql(t table) {
	var sqlStatement strings.Builder

	sqlStatement.WriteString("CREATE TABLE IF NOT EXISTS " + t.name + "( ")

	for key, element := range t.columnsType {
		sqlStatement.WriteString(t.columnName[key] + " " + element + ", ")
	}
	fmt.Println(sqlStatement.String())
}
