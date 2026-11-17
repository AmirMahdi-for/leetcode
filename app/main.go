package main

import (
    "fmt"
    "murph.com/app/product"
)

func main() {
	p := product.GetProduct("iphone13", 300.40, 29)

	fmt.Println(p)
	fmt.Println(p.IsAvailable())
}
