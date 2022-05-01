package main

import (
	"database/sql"

	dpAdapter "github.com/izabelrodrigues/hexagonal-fc/adapters/db"
	"github.com/izabelrodrigues/hexagonal-fc/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := dpAdapter.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)

}
