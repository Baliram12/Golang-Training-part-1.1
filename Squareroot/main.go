package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Print("Enter a number:")
	fmt.Scan(&num)
	result := math.Sqrt(num)
	fmt.Println("Squre root is:", result)
}
