package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	//e1
	e := Employee{}
	fmt.Printf("%v\n", e)
	//e2 con constructor
	e2 := Employee{
		id:       1,
		name:     "Name 2",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)
	//e3 equivalente a 1
	e3 := new(Employee) //Devuelve un apuntador a la instancia de Employee
	e3.id = 1
	e3.name = "Name 3"
	fmt.Printf("%v\n", *e3)
	//e4
	e4 := NewEmployee(1, "Name 4", false)
	fmt.Printf("%v\n", *e4)
}
