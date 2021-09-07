package main

import "fmt"

//función variadica
func sum(values ...int) int {
	total := 0

	for _, num := range values {
		total += num
	}

	return total
}

//retornos con nombres
func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(getValues(2))
}
