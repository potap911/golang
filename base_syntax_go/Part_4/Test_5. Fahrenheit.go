package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var F float64
    fmt.Scanf("%f", &F)

    var C float64 = (F - 32) * 5/9

    fmt.Println(C)    
}
