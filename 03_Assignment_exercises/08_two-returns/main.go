// Jacob Rachal
// 2-18-16
// Fresno State CSCI 130
// 8. Write a function which takes an integer.
// The function will have two returns.
// The first return should be the argument divided by 2.
// The second return should be a bool that let’s us know whether or not the argument was even.
// For example:
//a. half(1) returns (0, false)
//b. half(2) returns (1, true)

package main //_8_two_returns

import "fmt"

func half(n int) (int, bool){
	return n / 2, n%2 == 0
}

func main() {
	h, even := half(15)
	fmt.Println(h, even)
	k, even2 :=  half(10)
	fmt.Println(k, even2)
	k, even3 :=  half(4)
	fmt.Println(k, even3)
}