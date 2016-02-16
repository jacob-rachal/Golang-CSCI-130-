// Jacob Rachal
// 2-16-16 Tuesday
// Fresno State CSCI 130
// 5. Print all of the even numbers between 0 and 100.
// Looping! Awesome!
package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++{
		if i%2 == 0{fmt.Println(i)}
	}
}
