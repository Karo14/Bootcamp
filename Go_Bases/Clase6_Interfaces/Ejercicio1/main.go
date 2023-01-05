package main

import "fmt"


func main(){
	Alumno1 := Alumnos{
		Nombre: "Angie",
		Apellido: "Carrillo",
		DNI: "1000",
		Fecha: "27/12/2022",
	}

	Alumno1.detalle()
}

type Alumnos struct{
	Nombre string
	Apellido string
	DNI string
	Fecha string
}

func (a Alumnos) detalle () {
	fmt.Printf("Nombre: %s\n", a.Nombre)
	fmt.Printf("Apellido: %s\n", a.Apellido)
	fmt.Printf("DNI: %s\n", a.DNI)
	fmt.Printf("Fecha: %s\n", a.Fecha)
}
