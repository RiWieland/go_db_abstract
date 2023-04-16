package main

type data_reader interface {
	reader()
}

type customer struct {
	id          int
	name        string
	age         int
	email       string
	adress      string
	phonenumber string
}
