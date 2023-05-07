package main

type table struct {
	// types included + keys
	name        string
	columnName  []string
	columnsType []string
	//primaryKey string
	// not null
}

type view struct {
	columns []interface{}
}
