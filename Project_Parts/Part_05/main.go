// Jacob Rachal 3-22-16
// Fresno State CSCI 130
// Project Step 5
// INSTR:  continuing to build our application,
// integrate HMAC into our application to ensure that nobody tampers with the cookie.
package main//Part_05
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
	//"strings"

)
type user struct{
	Name string //split into two? Fname and Lname
	Age string
	//HMAC string
	//uuid *uuid.UUID //Great. I can't even remember why I had these here.
}
func snackWells( res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("temp.html")
	if err != nil{
		log.Fatalln("Something went wrong: ", err)
	}
	//name:= req.FormValue("Name:")
	//age:= req.FormValue("Age:")
	x := user{
		Name: req.FormValue("Name"),
		Age: req.FormValue("Age"), // is having a comma correct?
	}
	b, err := json.Marshal(x)
	if err != nil{
		fmt.Println("Error: ", err)//log.Fatalln("Error:", err) , or should it be Fprintf?
	}
	y :=base64.StdEncoding.EncodeToString(b)
	//new cookie if needed
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-fino",
			Value: id.String()+"|"+getCode(id.String()),//name+age,
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	cookie.Value = cookie.Value + "|" + y + "|" + getCode(y)//cookie.Value = y
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
}// above function from 035_sessions/05_HMAC/main.go
func main(){
	//code := getCode("test@example.com")
	//fmt.Println(code)//this and above from 035_sessions/05_HMAC/main.go
	http.HandleFunc("/", snackWells)
	/*http.HandleFunc("/", foo)*/
	http.Handle("/favicon.ico", http.NotFoundHandler()) //needed?
	//log.Println("Listening...")
	http.ListenAndServe(":8080", nil)

}
