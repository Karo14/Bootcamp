package main

import "fmt"

var sueldo float64
var Notas []float64
var Nota float64
var Categoria string
var Minutos int

func main() {
	fmt.Printf("Ingrese el sueldo: ")
	fmt.Scanln(&sueldo)
	fmt.Printf("\n")
	fmt.Printf("El sueldo es: %f \n", CalcularSalario(sueldo))

	fmt.Printf("Ingrese las notas: ")
	for i:=0; i<5; i++ { 
		fmt.Scanln(&Nota)
		Notas = append(Notas,Nota)
	}
    fmt.Printf("\n")
	fmt.Printf("El promedio es: %f \n", CalcularPromedio(Notas))

	fmt.Printf("Ingrese los minutos trabajados: ")
	fmt.Scanln(&Minutos)
	fmt.Printf("\n")

	fmt.Printf("Ingrese la categoria: ")
	fmt.Scanln(&Categoria)
	fmt.Printf("\n")
	
	fmt.Printf("El sueldo por categoria es: %f \n", CalcularSalarioXCategoria(Minutos,Categoria))
}

func CalcularSalario(salario float64) (sueldo float64){
	var impuesto1 = 0.17
	var impuesto2 = 0.27
	if(salario>150000){
		sueldo = salario - (salario * impuesto2)
		return sueldo
	}else if(salario>50000){
		sueldo = salario - (salario * impuesto1)
		return sueldo
	} else {
		return salario
	}
}

func CalcularPromedio(notas []float64) (promedio float64){
	var SumaTotal float64 = 0
	for _, num := range notas{
		if(num < 0){
			return 0
		}
		SumaTotal += num
	}
	promedio = float64(SumaTotal) / float64(len(notas))
	return promedio
}

func CalcularSalarioXCategoria(minutostrabajados int, categoria string)float64{
	var horas = minutostrabajados / 60
	switch categoria{
	case "A":
		return 3000 * float64(horas) + (1500)
	case "B":
		return 1500 * float64(horas) + (300)
	case "C":
		return 1000 * float64(horas)
	}
	return 0
}