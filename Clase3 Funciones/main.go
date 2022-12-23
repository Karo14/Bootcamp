package main

import "fmt"

func main() {
	fmt.Printf("El sueldo es: %f \n", CalcularSalario(1500000))
	fmt.Printf("El promedio es: %f \n", CalcularPromedio(3200,100,46,128,701))
	fmt.Printf("El sueldo por categoria es: %f \n", CalcularSalarioXCategoria(3600,"A"))
}

func CalcularSalario(salario float64) (sueldo float64){
	if(salario>50000){
		sueldo = salario - (salario * 0.17)
		if(salario>150000){
			sueldo = sueldo - (sueldo * 0.10)
		}
	}
	return sueldo
}

func CalcularPromedio(numeros ...int) (promedio float64){
	var SumaTotal = 0
	for _, num := range numeros{
		SumaTotal += num
	}
	promedio = float64(SumaTotal) / float64(len(numeros))
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