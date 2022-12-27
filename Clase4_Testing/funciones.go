package main

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