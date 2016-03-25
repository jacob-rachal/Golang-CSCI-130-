// Jacob Rachal 3-20-16
// Fresno State CSCI 130
// INSTR: Create a webpage which writes a cookie to the client's machine.
// Though this is NOT A BEST PRACTICE, you will store some session data in the cookie.
// Make sure you use HMAC to ensure that session data is not changed by a user.
package main//_3_cookie_writer_HMAC
import(
	"fmt"
	"io"
	"crypto/hmac"
	"crypto/sha256"
	//"github.com/nu7hatch/gouuid"
	"net/http"
)
func getCode(data string) string { //making the HMAC
	h := hmac.New(sha256.New, []byte("H3110w0rld")) //was previously "ourkey"
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
/*func cookieID(res http.ResponseWriter, req *http.Request){
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
	//fmt.Fprint(res, "Name: ", cookie.Name, "\nValue: ", cookie.Value, "\nHttpOnly: ", cookie.HttpOnly)
}*/
func foo(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	cookie, err := req.Cookie("session-ID")
	if err != nil {
		cookie = &http.Cookie{
			Name: "session-ID",
			Value: "Cookie-Value",//id.String(),
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	//var code string

	if req.FormValue("email") != "" {
		email := req.FormValue("email")
		cookie.Value = email + `|` + getCode(email)
	}

	io.WriteString(res, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	    ` + cookie.Value + `
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	  </body>
	</html>`)
}
func main(){
	//http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}