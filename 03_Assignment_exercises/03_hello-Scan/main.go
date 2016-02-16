// Jacob Rachal
// 2-14-16
// Fresno State CSCI 130
// 3. Create a program that prints to the terminal asking for a user to enter their name.
// Use fmt.Scan to read a user’s name entered at the terminal.
// Print “Hello <NAME>” with <NAME> replaced with what the user entered at the terminal.
package main

import "fmt"

func main (){
	var name string
	fmt.Println("Please enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)

}
