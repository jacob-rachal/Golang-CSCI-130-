// Jacob Rachal
// 2-18-16 Thursday
// Fresno State CSCI 130
// 6. Write a program that prints the numbers from 1 to 100.
// But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz".
// For numbers which are multiples of both three and five print "FizzBuzz".

package main //_6_fizz_buzz

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, " -- FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, " -- Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, " -- Buzz")
		} else {
			fmt.Println(i)
		}
	}
}