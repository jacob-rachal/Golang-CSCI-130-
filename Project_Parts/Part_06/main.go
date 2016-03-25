// Jacob Rachal 3-24-16
// Fresno State CSCI 130
// Project Step 6
// INSTR: create a page which illustrates what happens when a user changes a cookie.
// You can hard-code a changed cookie into your code.
package main//Part_06
import(
	"net/http"
	"html/template"
	"log"
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
	Name string //split into two? Fname and Lname
	Age string
	//HMAC string
	//uuid *uuid.UUID
}
func snackWells( res http.ResponseWriter, req *http.Request) { // I really need to break this up into smaller parts.
	tpl, err := template.ParseFiles("temp6.html")
	//NOTE: should I have that part in the form that says <h1>{{.FirstName}} {{.LastName}}</h1> ?
	if err != nil{
		log.Fatalln("Something went wrong: ", err)
	}
	//name:= req.FormValue("Name:")
	//age:= req.FormValue("Age:")
	x := user{
		Name: req.FormValue("Name"),
		Age: req.FormValue("Age"),
	}
	b, err := json.Marshal(x)
	if err != nil{
		fmt.Println("Error: ", err)//log.Fatalln("Error:", err) , or should it be Printf?
	}
	y :=base64.StdEncoding.EncodeToString(b)
	//new cookie if needed
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-fino",
			Value: id.String()+"|"+getCode(id.String()) +"|"+y+"|"+getCode(y),//name+age,
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	//checking done here?
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	cv := strings.Split(cookie.Value, "|")
	orgUuid := cv[0]
	orgCodeUuid := cv[1]
	if (getCode(orgUuid) == orgCodeUuid) {
		fmt.Fprintf(res, "Great News! The uuid has not been tampered with! \n")
	} else {
		fmt.Fprintf(res,"WARNING! The uuid has been changed! Go to battle stations! \n")
	}

	// User changes uuid.
	temp := cv[0]+"123"
	for i:=1; i < len(cv); i++ {
		temp = temp + "|" + cv[i]
	}
	cookie.Value = temp

	// cookie would be reset at this time (but is not done here).

	dv := strings.Split(cookie.Value, "|")
	newUuid := dv[0]
	newCodeUuid := dv[1]
	if (getCode(newUuid) == newCodeUuid) {
		fmt.Fprintf(res, "Good News! The uuid has not been tampered with! \n")
	} else {
		fmt.Fprintf(res,"WARNING! The uuid has been changed! Abort! Abort! Abort! \n")
	}

	//cookie.Value = cookie.Value + "|" + y + "|" + getCode(y)//cookie.Value = y
	http.SetCookie(res, cookie)
	err = tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func getCode(data string) string { //making the HMAC
	h := hmac.New(sha256.New, []byte("H3110w0rld")) //was previously "ourkey"
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func main(){
	//code := getCode("test@example.com")
	//fmt.Println(code)//this and above from 035_sessions/05_HMAC/main.go
	http.Handle("/favicon.ico", http.NotFoundHandler()) //needed?
	http.HandleFunc("/", snackWells)
	/*http.HandleFunc("/", foo)*/
	//log.Println("Listening...")
	http.ListenAndServe(":8080", nil)

}