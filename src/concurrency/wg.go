package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//WaitGroup
	//Un contador que está pendiente, que suma cada vez que nosotros creamos una goroutine
	//y resta cuando una goroutine ha terminado
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	//Bloquea el programa hasta que el contador quede en cero.
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {
	//defer siempre se ejecuta al final de cada funcion
	defer wg.Done()
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}
