package main

type Product struct {
	Name  string
	Price float64
	Stock int
}

func getProduct(name string, price float64, stock int) Product {
	return Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}
}
