package product

type Product struct {
	Name  string
	Price float64
	Stock int
}

func GetProduct(name string, price float64, stock int) Product {
	return Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}
}

func (p Product) IsAvailable() bool {
    return p.Stock > 0
}