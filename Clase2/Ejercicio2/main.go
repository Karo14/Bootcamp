package main

import "fmt"

func main(){
	prestamo(23,2,105000)
}

func prestamo(edad int, antiguedad int,sueldo float64){
	fmt.Printf("\n")
	if(edad<=22){
		fmt.Println("La edad requerida para este prestamo es mayor a 22 años de edad")
	} else if(antiguedad<=1){
		fmt.Println("Debe tener mas de 1 año de antiguedad para adquirir el prestamo")
	} else if(sueldo<=100000){
		fmt.Println("Se le cobrará interés del prestamo")
	} else{
		fmt.Println("Prestamo aprobado!!")
	}
	fmt.Printf("\n")
}