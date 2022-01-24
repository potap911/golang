package main

import "fmt"

func main() {
	fmt.Println(theFunc(4))

}

func theFunc(x int) (int, bool) {
	a := 0
	if x%2 == 0 {
		a = x / 2
		return a, true
	}
	return a, false
}
