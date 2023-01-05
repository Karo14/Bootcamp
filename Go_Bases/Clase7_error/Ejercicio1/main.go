package main

import "fmt"

type error interface{
	error() string
}

type Error struct{
	Error string
}

func main(){
	var salario int
	fmt.Print("Ingrese el salario: ")
	fmt.Scan(&salario)

	salary1 := Salary(salario)
	salary1.error()
	fmt.Println(salary1)

	fmt.Print("Ingrese el salario: ")
	fmt.Scan(&salario)

	salary2 := Salary(salario)
	salary2.error()
	fmt.Println(salary2)
}

func (e Error) error() string{
	return e.Error
}

func Salary(salary int) (metodo error){
	if salary < 150000 {
		metodo = Error{"Error: el salario ingresado no alcanza el mínimo imponible"} 
	} else {
		metodo = Error{"Debe pagar impuesto"}
	}
	return metodo
}