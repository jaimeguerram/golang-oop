package main

import (
	"fmt"
	"sync"
	"time"
)

//Buffered channels como semáforos
func main() {
	//Limita la capacidad de ejecutar 2 goroutines al mismo tiempo
	c := make(chan int, 2) //un canal para transmitir data de tipo entero
	var wg sync.WaitGroup  // Para sincronizar nuestras goroutines
	for i := 0; i < 10; i++ {
		c <- 1 //Se le da un cupo para ejecutar go doSomething(), bloquea el programa hasta que el canal tenga un cupo para ejecutar una goroutine
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done() //decrementa en uno nuestro contador
	fmt.Printf("Id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d finished\n", i)

	<-c //una vez que termina doSomthing libera el canal
}
