package main

type product struct {
	name, price string
}

func (p product) Name() string {
	return p.name
}

func (p product) Price() string {
	return p.price
}
