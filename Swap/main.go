package main

import "fmt"

func main() {
	fmt.Println(myfunction())
}

func myfunction() []int {
	a, b := 10, 5
	b, a = a, b
	return []int{a, b}
}
