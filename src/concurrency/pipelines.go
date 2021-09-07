package main

import "fmt"

// Generator recibe un canal de escritura (chan<-)
func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i //enviamos el valor al canal.
	}
	// se cerrará el canal y desbloqueará cualquier cosa que esté esperando
	// algún tipo de respuesta
	close(c)
}

// Double recibe un canal de entrada(lectura) y de salida(escritura) de enteros
func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

// Print recibe un canal de lectura
func Print(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

//Definiendo channels de lectura y escritura
//Cuando se trabaja con channels existe la gran probabilidad de crear
//un deadkicj si no somos cuidadosos con su utilización, una forma de
//mitigar parte de este riesgo es definiendo canales de lectura y escritura,
//pero no ambos.
func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}
