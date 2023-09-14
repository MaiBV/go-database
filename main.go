package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // o underline serve para não remover a dependencia do driver que não esta sendo usada diretamente
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	product1 := NewProduct("Arroz", 30.50)

	InsertProduct(db, *product1)

}

func InsertProduct(db *sql.DB, product Product) error {

	cmd, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?,?,?)")

	if err != nil {
		return err
	}

	defer cmd.Close()

	_, err = cmd.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}
