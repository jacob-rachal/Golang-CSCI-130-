// Jacob Rachal 3-20-16
// Fresno State CSCI 130
// INSTR: Create a webpage which writes a cookie to the client's machine.
// This cookie should be designed to create a session and
// should use a UUID, HttpOnly, and Secure (though you'll need to comment secure out).
package main//_2_cookie_writer
import(
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)
func cookieID(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session-ID")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-ID",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Fprint(res, "Name: ", cookie.Name, "\nValue: ", cookie.Value, "\nHttpOnly: ", cookie.HttpOnly)
	// Wait a minute what's going on here?!
	// The first time the page is loaded everything checks out fine, but when reloaded the HttpOnly then says false.
}
func main(){
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", cookieID)
	http.ListenAndServe(":8080", nil)
}