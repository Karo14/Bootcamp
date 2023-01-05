package main

import "fmt"

func main(){
	var salario int
	fmt.Print("Ingrese el salario: ")
	fmt.Scan(&salario)

	err1 := Salary(salario)
	fmt.Println(err1)

	fmt.Print("Ingrese el salario: ")
	fmt.Scan(&salario)

	err2 := Salary(salario)
	fmt.Println(err2)
}

func Salary(salary int) error{
	if salary < 150000 {
		return fmt.Errorf("Error: el salario ingresado no alcanza el mínimo imponible")
	} else {
		return fmt.Errorf("Debe pagar impuesto")
	}
}