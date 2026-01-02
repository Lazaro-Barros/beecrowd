package main

import "fmt"

const (
	pi float64 = 3.14159
)

func main() {
	var input float64
	fmt.Scan(&input)
	fmt.Printf("A=%.4f\n", pi*(input*input))
}
