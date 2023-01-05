package main

import (
	"errors"
	"fmt"
)
type error interface {
	Error() string
}

type MensajeError struct {
	errores int
}

func (e MensajeError) Error() string {
	return fmt.Sprintf("Ha ocurrido un error! %v",e.errores)
}

func validarSalary(salary int) error{

	if salary < 10000 {
		return MensajeError{101}
	} else {
		return nil
	}
}

func main(){
	var salario int
	fmt.Print("Ingrese el salario: ")
	fmt.Scan(&salario)

	err := validarSalary(salario)

	// La validación con la función “Is()”.
	if errors.Is(err,MensajeError{}) {
		fmt.Println(err)
		return
	}
}


