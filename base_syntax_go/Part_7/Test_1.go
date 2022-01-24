package main

import "fmt"

func main() {
	x := []float64{43, 40, 59, 10, 17}
	fmt.Println(sum(x))
}

func sum(x []float64) float64 {
	total := 0.0
	for _, v := range x {
		total += v
	}

	return total
}
