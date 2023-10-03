package models

import "alura/go-webshop-app/database"

type Product struct {
	Id                int32
	Name, Description string
	Price             float64
	Quantity          int32
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
		var id, quantity int32
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
