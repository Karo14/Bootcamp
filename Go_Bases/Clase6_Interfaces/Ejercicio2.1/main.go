package main

import "fmt"

const (
	Producto_Pequeño = "P"
	Producto_Mediano = "M"
	Producto_Grande = "G"
)

type Producto struct{
	Tipo_producto string
	Precio float64
}

func main(){
	Producto1 := Producto{
		Tipo_producto: "P",
		Precio: 4500,
	}
	fmt.Printf("El precio base: %f\n",Producto1.precio())
	fmt.Println(factory("P",5600))
}

func (p Producto)precio()float64{
	if(p.Tipo_producto == Producto_Pequeño){
		return p.Precio
	} else if(p.Tipo_producto == Producto_Mediano){
		return p.Precio + (p.Precio * 0.03)
	} else if(p.Tipo_producto == Producto_Grande){
		return p.Precio + (p.Precio * 0.06) + 2500
	}
	panic("Producto Incorrecto!!")
}

func (p Producto) String() string{
	return fmt.Sprintf("%s,%f",
		p.Tipo_producto,
		p.Precio)
}

func factory(tipo_producto string, precio float64) string{
	var llenarproducto fmt.Stringer = Producto{
										Tipo_producto: "P",
										Precio: 4500,
									}
	return llenarproducto.String()
}
