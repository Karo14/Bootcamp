package main

import (
	"errors"
	"fmt"
)

func main(){
	var salario int
	fmt.Print("Ingrese el salario: ")
	fmt.Scan(&salario)

	err := validarSalary(salario)

	if(err != nil) {
		fmt.Println(err)
	}

}
 
func validarSalary(salary int) error{

	if salary < 10000 {
		return errors.New("Error: el salario es menor a 10.000")
	} else {
		return nil
	}
}