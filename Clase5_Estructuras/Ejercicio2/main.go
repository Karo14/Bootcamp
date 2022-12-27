package main

import "fmt"

type Person struct{
	ID int
	Name string
	DateOfBirth string
}

type Employee struct{
	ID int
	Position string
	Person Person
}

func (e Employee)PrintEmployee(Employee){
	fmt.Println("Name: ",e.Person.Name)
	fmt.Println("DateOfBirth: ",e.Person.DateOfBirth)
	fmt.Println("Position: ",e.Position)
}

func main(){
	fmt.Printf("\n")
	p := Person{
		ID: 1,
		Name: "Angie Carrillo",
		DateOfBirth: "10/05/1999",
	}

	e := Employee{
		ID: 1,
		Position: "Software Developer",
		Person: p,
	}

	e.PrintEmployee(e)

	fmt.Printf("\n")
}