package main

import "fmt"

func main() {
	fmt.Println(theFunc(1, 2, 3))

}

func theFunc(x ...int) int {
	max := 0
	for i := 0; i < len(x); i++ {
		if max < x[i] {
			max = x[i]
		}
		return max
	}
}
