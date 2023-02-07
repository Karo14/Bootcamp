package main

import (
	"PruebaBD_Warehouse/internal/domain"
	"PruebaBD_Warehouse/internal/products"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Para instalar go get github.com/go-sql-driver/mysql
func main (){
	databaseConfig := mysql.Config{
		User: "root",
		Passwd: "",
		Addr: "localhost:3306",
		DBName: "my_db",
		ParseTime: true,
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

	// //Method STORE
	tm,_ := time.Parse("2006-01-02","2023-01-26")
	s := domain.Product{
		Name : "Blusa",
		Quantity : 5,
		Code_value : "05",
		Is_published: "NO",
		Expiration: tm,
		Price : 15000,
	}

	err = repository.Store(&s)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Insertado con Exito!")
	
	//Method UPDATE
	tm1,_ := time.Parse("2006-01-02","2023-01-26")
	p := domain.Product{
		ID: 201,
		Name : "Blusa",
		Quantity : 6,
		Code_value : "05",
		Is_published: "1",
		Expiration: tm1,
		Price : 16,
	}

	err = repository.Update(&p)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Actualizado con Exito!")

	// // //Method DELETE
	err = repository.Delete(199)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Eliminado con Exito!")
}