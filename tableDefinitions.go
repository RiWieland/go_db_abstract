package main

type customerCollection struct {
	table
	c []customer
}

type orderCollection struct {
	table
	o []order
}

type customer struct {
	id        int
	firstname string
	lastname  string
	age       int
	adress
}

type order struct {
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
