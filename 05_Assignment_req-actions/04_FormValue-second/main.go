// Jacob Rachal 3-13-16
// Fresno State CSCI 130
// INSTR: Create a webpage that serves a form and allows the user to enter their name.
// Once a user has entered their name, show their name on the webpage. Use req.FormValue to do this
package main//_4_FormValue_second
import(
	"net/http"
	"io"
)
func main(){
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func( res http.ResponseWriter, req *http.Request){
		val := req.FormValue("name")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<form method="POST">
			Enter your name please:
		 	<input type="text" name="name">
		 	<input type="submit">
    			</form>`+val)
	})
	http.ListenAndServe(":8080", nil)
} //I really hope this is what is being asked for.