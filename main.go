package main

import "fmt"

func main() {
	var x float64

	fmt.Scan(&x)

	e := 0.0
	fact := 1.0
	n := 0.0
	for n <= x {
		e += 1 / (fact)
		n++
		fact *= n
	}
	fmt.Println(e)
}
