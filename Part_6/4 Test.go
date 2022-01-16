package main

import "fmt"

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 8, 83, 27, 19, 97, 9, 17}
	low := x[0]
	for i := 0; i < len(x); i++ {
		if low > x[i] {
			low = x[i]
		}
	}
	fmt.Println(low)
}
