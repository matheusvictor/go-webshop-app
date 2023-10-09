package models

import (
	"alura/go-webshop-app/database"
)

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func FindAll() []Product {
	db := database.ConnectDatabase()

	selectAllProducts, err := db.Query("SELECT * FROM products ORDER BY name ASC")
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

		p.Id = id
		p.Name = name
		p.Quantity = quantity
		p.Description = description
		p.Price = price

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func FindById(id int) Product {
	db := database.ConnectDatabase()

	selectedProduct, err := db.Query("SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	foundedProduct := Product{}

	for selectedProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := selectedProduct.Scan(
			&id,
			&name,
			&description,
			&quantity,
			&price,
		)
		if err != nil {
			panic(err.Error())
		}

		foundedProduct.Id = id
		foundedProduct.Name = name
		foundedProduct.Quantity = quantity
		foundedProduct.Description = description
		foundedProduct.Price = price
	}
	defer db.Close()
	return foundedProduct
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

func Delete(id int) {
	db := database.ConnectDatabase()

	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}
	deleteProduct.Exec(id)
	defer db.Close()
}

func Update(id int, name string, desc string, quantity int, price float64) {
	db := database.ConnectDatabase()

	update, err := db.Prepare("UPDATE products SET name=$1, description=$2, quantity=$3, price=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(name, desc, quantity, price, id)
	defer db.Close()
}
