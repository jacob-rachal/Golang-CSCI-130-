// Jacob Rachal 3-9-16
// Fresno State CSCI 130
// INSTR: Create a webpage that serves at localhost:8080
// and will display the name in the url when the url is localhost:8080/name
// - use req.URL.Path to do this
package main//_2_display_name
/*import(
	"fmt"
	"net/http"
)
func main(){
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		res.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(res, req.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}*/
import (
	"io"
	"net/http"
)

type jojo int

func (h jojo) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<h1>`+req.URL.Path+`</h1><br>`)
}

func main() {
	var badger jojo

	mux := http.NewServeMux()
	mux.Handle("/", badger)

	http.ListenAndServe(":8080", mux)
}