package main

type customer struct {
	table
	id        int
	firstname string
	lastname  string
	age       int
	adress
}

type orders struct {
	table
	id        int
	firstname string
	lastname  string
	object    string
	amount    int
	shipped   bool
}

type adress struct {
	streetAddress string
	city          string
	state         string
}
