package main

import (
	//"PruebaBD/internal/products"
	"PruebaBD/internal/products"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// Para instalar go get github.com/go-sql-driver/mysql
func main (){
	databaseConfig := mysql.Config{
		User: "root",
		Passwd: "",
		Addr: "localhost:3306",
		DBName: "storage",
	}
	println(databaseConfig.FormatDSN())
	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil{
		panic(err)
	}

	err = database.Ping()
	if err != nil{
		panic(err)
	}

	println("Database conncection established")

	// Crear repositorio
	var repository products.Repository = &products.MySQLRepository{
		Database: database,
	}
	
	//Method GET
	product, err := repository.Get(1)
	if err!=nil {
		panic(err)
	}
	fmt.Println(product)

	//Method STORE
	s := products.Product{
		Name : "Blusa",
		Type : "Camisa",
		Count : 5,
		Price : 15000,
	}

	err = repository.Store(&s)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Insertado con Exito!")
	
	//Method UPDATE
	p := products.Product{
		ID: 2,
		Name : "Blusa",
		Type : "Camisa",
		Count : 50,
		Price : 15000,
	}

	err = repository.Update(&p)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Actualizado con Exito!")

	//Method DELETE
	err = repository.Delete(3)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Eliminado con Exito!")
}