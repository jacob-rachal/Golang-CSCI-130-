//Jacob Rachal 1/28/2016
//for CSCI 130 "Web Programing"
//ASSIGNMENT: Create a program that shows the type of some variable (use fmt.Printf).

package main

import "fmt"

func main() {
     x := 55
     y := "hello"
     fmt.Printf("The type of variable x is %T.	", x)
     fmt.Printf("The type of variable y is %T.", y)
}