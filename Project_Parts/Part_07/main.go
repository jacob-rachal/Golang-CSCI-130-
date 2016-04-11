// Jacob Rachal 3-24-16
// Fresno State CSCI 130
// Project Step 7
// INFO: PROJECT STEP 7 - Allow the user to login.
// Store the information about whether or not a user is logged-in
// in both the "user" data type you created and in the cookie.
// Show a "logout" button when the user is logged in
package main//Part_07
import(
	"net/http"
	"html/template"
	"log" //remove?
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"
)
type user struct{
	Name string
	Age string
	LogStatus bool
}
var tpl *template.Template

/*func init() {
	tpl, _ = template.ParseGlob("templates/*.html")
}*/

func index(res http.ResponseWriter, req *http.Request) {

	cookie := genCookie(res, req)
	u := decodeUser(cookie)

	if req.Method == "POST" {
		//src, hdr, err := req.FormFile("data")
		u.LogStatus = true
		u.Name = req.FormValue("name")
		u.Age = req.FormValue("age")

		xs := strings.Split(cookie.Value, "|")
		id := xs[0]

		cookie = currentVisitor(u, id)
		http.SetCookie(res, cookie)

	}
	//m := Model(cookie)
	fmt.Println("DEBUG: 52  ",u) //just for debugging
	tpl.ExecuteTemplate(res, "temp7.html", u)
}
/* body of temp7 temporairly stored here for debugging.
{{if .LogStatus}}
    <h1> The Concordiant Template </h1>
    <h2>By Jacob Rachal</h2>
    <h3>Your Page</h3>
        <p>{{.Name}}</p><br>
        <p>{{.Age}}</p>

    <br>
    <a href="/logout">Log Out</a>
    {{end}}
    <br>
{{else}}
    <a href="/login">Please Login</a>
{{end}}
*/
func login(res http.ResponseWriter, req *http.Request) {

	cookie := genCookie(res, req)

	if req.Method == "POST" && req.FormValue("passwd") == "secret" {
		u := decodeUser(cookie)
		u.LogStatus = true
		u.Name = req.FormValue("usernm")

		xs := strings.Split(cookie.Value, "|")
		id := xs[0]

		cookie := currentVisitor(u, id)
		http.SetCookie(res, cookie)

		http.Redirect(res, req, "/", 302)
		return
	}
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	cookie := newVisitor()
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", 302)
	/*`
	<form method="POST">
	    <p> Logout things here </p>

	</form>
	`*/
}

func genCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {

	cookie, err := req.Cookie("session-id")
	if err != nil {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	if strings.Count(cookie.Value, "|") != 2 {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	if tampered(cookie.Value) {
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	return cookie
}
func decodeUser(c *http.Cookie) user { //might be known as the Model.go file
	xs := strings.Split(c.Value, "|")
	usrData := xs[1]

	bs, err := base64.URLEncoding.DecodeString(usrData)
	if err != nil {
		log.Println("Error decoding base64", err)
	}

	var u user
	err = json.Unmarshal(bs, &u)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
		return user{}
	}
	return u
}

func newVisitor() *http.Cookie {
	mm := initialUser()
	id, _ := uuid.NewV4()
	return makeCookie(mm, id.String())
}

func currentVisitor(u user, id string) *http.Cookie {
	mm := marshalUser(u)
	return makeCookie(mm, id)
}

func makeCookie(mm []byte, id string) *http.Cookie {
	b64 := base64.URLEncoding.EncodeToString(mm)
	code := getCode(b64)
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}

func marshalUser(u user) []byte {
	bs, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return bs
}

func initialUser() []byte {
	u := user{
		Name:   "",
		Age:  "",
		LogStatus: false,
	}
	return marshalUser(u)
}
func main(){
	tpl, _ = template.ParseGlob("templates/*.html")

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}
func getCode(data string) string { //making the HMAC
	h := hmac.New(sha256.New, []byte("H3110w0rld")) //was previously "ourkey"
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func tampered(s string) bool {
	xs := strings.Split(s, "|")
	usrData := xs[1]
	usrCode := xs[2]
	if usrCode != getCode(usrData) {
		return true
	}
	return false
}
