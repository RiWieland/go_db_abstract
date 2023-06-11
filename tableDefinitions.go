package main

type customerTable struct {
	C []customer
}

type orderTable struct {
	C []order
}

type customer struct {
	Id        int
	Firstname string
	Lastname  string
	Age       int
	Adress
}

type order struct {
	Id        int
	Firstname string
	Lastname  string
	Object    string
	Amount    int
	Shipped   bool
}

type Adress struct {
	streetAddress string
	city          string
	state         string
}
