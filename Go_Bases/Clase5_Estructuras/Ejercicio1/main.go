package main

import "fmt"

type Product struct{
	ID int 
	Name string
	Price float64 
	Description string
	Category string
}

var Products []*Product

func (p *Product) Save(Product) {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Println(*product)
	}
}

func getByld(ID int) *Product {
	for _, prod := range Products {
		if prod.ID == ID {
			return prod
		}
	}
    panic("Articulo no econtrado")
}
func main(){
	fmt.Printf("\n")
	Producto1 := Product{ID: 1, Name: "Camisa", Price: 15000, Description: "Camisa Blanca de cuello", Category: "Camisas"}
	Producto1.Save(Producto1)

	Producto2 := Product{ID: 2, Name: "Pantalon", Price: 25000, Description: "Pantalon Jean", Category: "Pantalon"}
	Producto2.Save(Producto2)

	Producto3 := Product{ID: 3, Name: "Tennis", Price: 30000, Description: "Tennis", Category: "Zapatos"}
	Producto3.Save(Producto3)

	fmt.Println("Consultar Productos: ")
	fmt.Printf("\n")

	Producto2.GetAll()

	fmt.Printf("\n")
	fmt.Println("Buscar Productos: ")
	fmt.Printf("\n")

	fmt.Println(getByld(1))
	fmt.Println(getByld(2))
	fmt.Println(getByld(3))
	fmt.Println(getByld(4))

	fmt.Printf("\n")
}