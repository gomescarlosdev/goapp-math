package main

import "fmt"

func main() {
	fmt.Println(Sum(10, 10))
	fmt.Println(Multi(10, 10))
}

func Sum(a int, b int) int {
	return a + b
}

func Multi(a int, b int) int {
	return a * b
}

func Sub(a int, b int) int {
	return a - b
}

func SumX(a int, b int) int {
	return (a + b) * a
}

func SubX(a int, b int) int {
	return (a * b) - a
}

func MultiX(a int, b int) int {
	return (a * b) - a
}
