// Jacob Rachal
// 2-25-16
// Fresno State CSCI 130
// 2. Modify the previous program to use a func expression.
/* */
package main //_9_two_returns_func
import "fmt"

func main(){
	half := func (n int) (int, bool){
		return n/2, n%2==0
	} //If I am understanding this right, 'func' turned 'half' into a function that can now be called. Huh.
	fmt.Println(half(15))
	fmt.Println(half(10))
	fmt.Println(half(4))
	// This returns the same result with about half the code. Nice!
}