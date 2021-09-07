package main

import "fmt"

//Los workers pools permiten la creación de múltiples trabajadores
//que llevarán a cabo determinadas tareas, en este caso Go puede
//explotar las Goroutines para alcanzar worker pools concurrentes
func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	//Se ejecutan los 3 workers
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	//Lee los resultados
	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

// Worker de los 3 workers el que esté escuchando se le asigna la tarea
func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
