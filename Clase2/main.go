package main

import "fmt"

var Nombre = "Angie"

func main() {
	Ejercicio1()
	Ejercicio2()
}

func Ejercicio1() {
	fmt.Println("Cantidad de caracteres:")
	fmt.Println(len(Nombre))
}

func Ejercicio2(){

	fmt.Println("Estas son las letras:")
	for _, letras := range Nombre{
		fmt.Println(string(letras))
	}
}

