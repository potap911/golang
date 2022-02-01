package main

import "fmt"

// Напишите программу, которая меняет местами два числа

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x

}
