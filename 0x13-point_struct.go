package main

import "fmt"

func main() {
	type Product struct {
		ProductName string
	}

	type InvoiceDetail struct {
		Id      int
		Product *Product
	}

	type Invoice struct {
		Details    *InvoiceDetail
		DetailsIds []int
	}

	var z Product
	var w InvoiceDetail
	var u Invoice
	u.Details = &w
	w.Product = &z
	u.Details.Product.ProductName = "Soap"

	fmt.Println(z)
	fmt.Println(w)
	fmt.Println(u)
}
