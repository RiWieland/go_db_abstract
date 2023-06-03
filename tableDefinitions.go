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
	Id        int
	Firstname string
	Lastname  string
	Object    string
	Amount    int
	Shipped   bool
}

type adress struct {
	streetAddress string
	city          string
	state         string
}
