// Jacob Rachal 3-24-16
// Fresno State CSCI 130
// Project Step 8
// INFO: PROJECT STEP 8 -  Allow the user to logout.
// Show a log-in button when the user is not logged-in.
// Show a log-out button only when the user is logged in.
package main//Part_08
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

func init() {
	tpl, _ = template.ParseGlob("templates/*.htmltemp")
}
func index(res http.ResponseWriter, req *http.Request) {

	cookie := genCookie(res, req)
	u := decodeUser(cookie)

	if req.Method == "POST" {
		u.LogStatus = true
		u.Name = req.FormValue("name")
		u.Age = req.FormValue("age")

		xs := strings.Split(cookie.Value, "|")
		id := xs[0]

		cookie = currentVisitor(u, id)
		http.SetCookie(res, cookie)

	}
	tpl.ExecuteTemplate(res, "index7.html", u)
}
func login(res http.ResponseWriter, req *http.Request) {

	cookie := genCookie(res, req)

	if req.Method == "POST" && req.FormValue("password") == "secret" {
		u := decodeUser(cookie)
		u.LogStatus = true
		u.Name = req.FormValue("name")

		xs := strings.Split(cookie.Value, "|")
		id := xs[0]

		cookie := currentVisitor(u, id)
		http.SetCookie(res, cookie)

		http.Redirect(res, req, "/", 302)
		return
	}
	tpl.ExecuteTemplate(res, "login.htmltemp", nil)
}
func logout(res http.ResponseWriter, req *http.Request) {
	cookie := newVisitor()
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", 302)
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
func decodeUser(c *http.Cookie) user {
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
