// Jacob Rachal 3-9-16
// Fresno State CSCI 130
// INSTR: Create a webpage that serves at localhost:8080
// and will display the name in the url when the url is localhost:8080/name
// - use req.URL.Path to do this
package main//_2_display_name
import(
	"fmt"
	"net/http"
)
func main(){
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		res.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(res, "Path: %v\n", req.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
/* While this is running, on the webpage you should see in the upper right corner "Path: /".
 The url address should say "localhost:8080/", and after the "/" you can type in anything you want
  (ex: "Hello_World"), hit enter or reload, and now the Path: should also show whatever phrase you added.
 */