package models

import "alura/go-webshop-app/database"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func FindAll() []Product {
	db := database.ConnectDatabase()

	selectAllProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := selectAllProducts.Scan(
			&id,
			&name,
			&description,
			&quantity,
			&price,
		)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Quantity = quantity
		p.Description = description
		p.Price = price

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func Create(name string, desc string, price float64, quant int) {
	db := database.ConnectDatabase()

	insertProduct, err := db.Prepare("INSERT INTO products (name, description, quantity, price) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertProduct.Exec(name, desc, quant, price)
	defer db.Close()
}
