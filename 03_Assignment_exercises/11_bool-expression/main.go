// Jacob Rachal 3-3-16
// Fresno State CSCI 130
// What's the value of this expression: (true && false) || (false && true) || !(false && false)?
package main //_1_bool_expression
import "fmt"
func main() {
	fmt.Println((true && false) || (false && true) || !(false && false))
} //true
