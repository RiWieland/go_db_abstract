package db

type table struct {
	// types included + keys
	name        string
	columnName  []string
	columnsType []string
	//primaryKey string
	// not null
}
