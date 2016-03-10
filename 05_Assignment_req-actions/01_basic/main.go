// Jacob Rachal 3-9-16
// Fresno State CSCI 130
// INSTR: Create a webpage that displays the URL path using req.URL.Path
//  The page should display / in the upper left corner of the webpage.
package main //_1_basic
import(
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		res.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(res, req.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
