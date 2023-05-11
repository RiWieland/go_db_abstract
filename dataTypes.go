package main

// Abstract type:

type column struct {
	name string
}

// Concrete types:
type typeInt struct {
	column
}

type typeString struct {
	column
}

type typeVarchar struct {
	column
}
