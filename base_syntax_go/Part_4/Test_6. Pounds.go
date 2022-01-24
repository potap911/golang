package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var P float64
    fmt.Scanf("%f", &P)

    M := P * 0.3048

    fmt.Println(M)    
}
