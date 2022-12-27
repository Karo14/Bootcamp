package main

import (
	"fmt"
	"encoding/json"
)

var nombre string
var apellido string
var edad int
var P2 Persona
var P3 Persona
var P4 Persona
var v int = 19
var p *int

type Persona struct{
	nombre string //`json:"nombre"`
	apellido string //`json:"apellido"`
	edad int //`json:"edad"`
	Gustos Preferencias
}

type Preferencias struct {
	Deporte string
	Hobbie string
}

func (p *Persona)SetNombre(nombre string) {
	p.nombre = nombre
}

func main(){
  fmt.Println("Coloque su Nombre: ")
  fmt.Scanln(&nombre)
  P3.nombre = nombre

  fmt.Println("Coloque su Apellido: ")
  fmt.Scanln(&apellido)
  P3.apellido = apellido

  fmt.Println("Coloque su Edad: ")
  fmt.Scanln(&edad)
  P3.edad = edad

  fmt.Println("Persona 3: ",P3)


  P1 := Persona{
	nombre: "Angie",
	apellido: "Carrillo",
	edad: 23,
  }

  P2.nombre = "Jonathan"
  P2.apellido = "Jimenez"
  P2.edad = 27

  fmt.Println("Persona 1 y 2: ",P1,P2)

  p4 := Persona{"Laura","Carrillo",17,Preferencias{"N/A","N/A"}}
  MiJson, err := json.Marshal(p4)

  fmt.Println("Persona 4: "+string(MiJson))
  fmt.Println(err)

  P4.SetNombre("Angie")
  fmt.Println("Persona 4: ",P4)

  p=&v
  fmt.Printf("El puntero p referencia a la direccion de memoria: %v \n",v) 
  fmt.Printf("Al desreferencia el puntero p obtengo el valor: %d \n",*p) 

}

