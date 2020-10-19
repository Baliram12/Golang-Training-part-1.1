package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	b, err := strconv.ParseFloat(os.Args[1], 10)

	if err != nil {
		fmt.Println("First input parameter must be of type float")
		os.Exit(1)
	}

	c, err := strconv.ParseFloat(os.Args[2], 10)

	if err != nil {
		fmt.Println("Second input parameter must be of type float")
		os.Exit(1)
	}

	fmt.Println("b : ", b)
	fmt.Println("c : ", c)

	// calculate the roots of the polynomial
	// x^2 + bx + c using the quadratic formula

	discriminant := b*b - 4.0*c
	d := math.Sqrt(discriminant)

	fmt.Printf("%0.6f\n", ((-b + d) / 2.0))
	fmt.Printf("%0.6f\n", ((-b - d) / 2.0))

}
