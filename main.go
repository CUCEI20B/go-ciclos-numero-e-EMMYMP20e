package main

import "fmt"

func main() {
	var e float64
	var n float64
	var fact float64

	var x float64

	fmt.Scan(&x)

	e = 0
	fact = 1
	n = 0
	for n < x {
		e += 1 / (fact)
		n++
		fact *= n
	}
	fmt.Println(e)
}
