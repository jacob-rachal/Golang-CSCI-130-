// Jacob Rachal
// 2-27-16
// Fresno State CSCI 130
// 10/3 write a function with one variadic parameter that finds the greatest number in a list of numbers.
package main//_0_variadic_parameter_find_greatest
import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}
func main(){
	fmt.Println("Running the function, what should be returned for this test is 1000.")
	greatest := max(1, 2, 4, 7, 11, 15, 1000, 20, 120, 59, 201, 477, 3, 809)
	fmt.Println(greatest)
}