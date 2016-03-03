// Jacob Rachal 3-3-16
// Fresno State CSCI 130
// program description here please (actually, how would I describe this program? Find out later.)
// parameters and arguments
package main //_2_foo_multicall
import "fmt"
func main(){
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
} //I guess this works?
func foo(numbers ...int){
	fmt.Println(numbers)
} //...this is, new.