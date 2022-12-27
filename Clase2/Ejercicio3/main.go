package main

import "fmt"

func main(){
	mes(12)
}

func mes(numero int){
	if(numero < 1 || numero > 12){
		fmt.Println("Numero de mes incorrecto")
	} else{
		if(numero == 1){
			fmt.Println("Mes de Enero")
		}

		if(numero == 2){
			fmt.Println("Mes de Febrero")
		}

		if(numero == 3){
			fmt.Println("Mes de Marzo")
		}

		if(numero == 4){
			fmt.Println("Mes de Abril")
		}

		if(numero == 5){
			fmt.Println("Mes de Mayo")
		}

		if(numero == 6){
			fmt.Println("Mes de Junio")
		}

		if(numero == 7){
			fmt.Println("Mes de Julio")
		}

		if(numero == 8){
			fmt.Println("Mes de Agosto")
		}

		if(numero == 9){
			fmt.Println("Mes de Septiembre")
		}

		if(numero == 10){
			fmt.Println("Mes de Octubre")
		}

		if(numero == 11){
			fmt.Println("Mes de Noviembre")
		}

		if(numero == 12){
			fmt.Println("Mes de Diciembre")
		}
	}
}