// Jacob Rachal 3-13-16
// Fresno State CSCI 130
// PROJECT PART 1 (aka STEP 1)
// INSTR: create a web application that serves an HTML template.
package main //Part_01
import(
	//"fmt"
	"net/http"
	"html/template"
	//"io"
	"log"
)

func handler( res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("temp.html")
	if err != nil{
		log.Fatalln("Something went wrong: ", err) //... was it we were told to NOT use this or that it was ok?
	}
	tpl.Execute(res, nil)
}
func main(){
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
// NOTES: I really, REALLY hope this is what is being asked for.