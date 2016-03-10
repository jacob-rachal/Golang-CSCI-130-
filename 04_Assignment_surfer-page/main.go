// Jacob Rachal 3-7-16
// Fresno State CSCI 130
package main //_4_Assignment_surfer_page
import (
	"net/http"
	"html/template"
	"log"
	//"fmt"
)

func surf(res http.ResponseWriter, req *http.Request) {
tpl, err := template.ParseFiles("templ.gohtml")
if err != nil {
log.Fatalln(err)
}

tpl.Execute(res, nil)
}
func main() {
	http.HandleFunc("/", surf)  // issue here. EDIT: fixed
	//favicon killer here
	//http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

http.ListenAndServe(":8080", nil)
}
