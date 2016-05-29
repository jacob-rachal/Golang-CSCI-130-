// Jacob Rachal 5/22/16 CSCI 130
// Create an application which demonstrates AJAX.
// One possible idea: an "is the name already taken" app. Users can enter words in a form field.
// The words will be stored in either memcache or the datastore. If the word is already stored,
// a message will display on the webpage letting them know that word is already stored.
// Here is some starting code for you. (example code is missing?)
package main//_4_Assignment_AJAX
import(
	"net/http"
	"fmt"
	"io/ioutil"
	"strconv"
)
var page string
var uniqueID int

func index(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, page)
}
func ajax(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","text/plain")
	fmt.Fprintf(res,`<div class="readable" onclick="ripAndtear('`+strconv.Itoa(uniqueID)+`')" id="`+ strconv.Itoa(uniqueID) +`">` + req.FormValue("key") + `</div>`)
	uniqueID++
}
func init(){
	temp, _ := ioutil.ReadFile("index.html")
	page = string(temp)
	uniqueID = 0
}
func main(){
	http.Handle("/files/",http.StripPrefix("/files",http.FileServer(http.Dir("files"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/ajax/", ajax)
	http.ListenAndServe(":8080",nil)
}