// Jacob Rachal 3-12-16
// Fresno State CSCI 130
// INSTR: Create a webpage which uses a cookie to track the number of visits of a user.
// Display the number of visits.
// Make sure that the favicon.ico requests are not also incrementing the number of visits.
package main//_1_tracking_visits
import(
	//"fmt"
	"io"
	"net/http"
	"strconv"
)
func bakery(res http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/" { //ignoring the favicon thing
		http.NotFound(res, req)
		return
	}
	//Aquire the cookie
	cookie, err := req.Cookie("THE-COOKIE")
	// If cookie is MIA, make a new one
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "THE-COOKIE",
			Value: "0",
		}
	}
	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(res, cookie)
	io.WriteString(res, cookie.Value)
}
func main(){
	http.HandleFunc("/", bakery)
	http.ListenAndServe(":8080", nil)
}