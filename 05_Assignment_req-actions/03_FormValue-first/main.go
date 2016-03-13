// Jacob Rachal 3-9-16  3-13-16
// Fresno State CSCI 130
// INSTR:Create a webpage that serves at localhost:8080 and will display
// the name in the url when the url is localhost:8080?n="some-name"
// - use req.FormValue to do this
package main//_3_FormValue_first
import(
	"net/http"
	"io"
)
func handler( res http.ResponseWriter, req *http.Request){
	key := "n" //Originally, I wasn't sure what to put here so I used "hello", but when I ran it and took another look at the instructions I changed it to "n"
	val := req.FormValue(key)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<form method="GET">
			 <input type="text" name="n">
			 <input type="submit">
	   		 </form>`+val)
} //...wait, am I also supposed to show the user input on the webpage as well? How do I not do that?
func main(){
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}