package main

import "fmt"

type Employee struct {
	id   int
	name string
}

//pointers receivers
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}
func main() {
	e := Employee{}

	e.SetId(5)
	e.SetName("Name")

	fmt.Println(e)
}
