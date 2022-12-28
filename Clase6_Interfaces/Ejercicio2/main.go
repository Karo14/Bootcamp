package main

import "fmt"

type Producto interface{
	precio() float64
}

type ProductoPequeño struct{
	preciop float64
}

type ProductoMediano struct{
	preciom float64
}

type ProductoGrande struct{
	preciog float64
}

func main(){
	Producto1 := factory("M", 2500)
	precio1 := Producto1.precio()
	fmt.Println("El precio es: ", precio1)

	Producto2 := factory("G", 2500)
	precio2 := Producto2.precio()
	fmt.Println("El precio es: ", precio2)
}

func (p ProductoPequeño) precio() float64{
	return p.preciop
}

func (p ProductoMediano) precio() float64{
	return p.preciom + (p.preciom * 0.03)
}

func (p ProductoGrande) precio() float64{
	return p.preciog + (p.preciog * 0.06) + 2500
}

func factory(tipo_producto string, precio float64) (metodo Producto){
	switch tipo_producto {
	case "P" : 
		metodo = ProductoPequeño{precio}
	case "M":
		metodo = ProductoMediano{precio}
	case "G":
		metodo = ProductoGrande{precio}
	}
	return metodo
}