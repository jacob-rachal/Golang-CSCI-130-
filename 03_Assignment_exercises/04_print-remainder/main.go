// Jacob Rachal
// 2-16-16 Tuesday
// Fresno State CSCI 130
// 4.Create a program that prints to the terminal asking for a user to enter a small number and a larger number.
// Print the remainder of the larger number divided by the smaller number.
// Note: The REMAINDER gets printed, not the quotient. Not x div y, but x % y.
package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&num1)
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&num2)
	fmt.Println(num1, "%", num2, " = ", num1%num2)
}