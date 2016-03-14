// Jacob Rachal 3-13-16
// Fresno State CSCI 130
// PROJECT PART 4 (aka STEP 4)
// INSTR: refactoring our application,
// create a new data type called "user" which has fields for the user's name and age.
// When you receive the user's name and age form submission,
// create a variable of type "user" then put those values from the form submission into the fields for that variable.
// Marshal your variable of type "user" to JSON. Encode that JSON to base64. Store that value in the cookie.
package main //Part_04
import(
	"net/http"
	"html/template"
	//"io"
	"log"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
)

type user struct{
	Name string
	Age string
}
func snackWells( res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("temp.html")
	if err != nil{
		log.Fatalln("Something went wrong: ", err)
	}
	name:= req.FormValue("Name:")
	age:= req.FormValue("Age:")
	x := user{
		Name: req.FormValue("Name:")
		Age: req.FormValue("Age")
	}
	b, err := json.Marshal(x)
	if err != nil{
		log.Fatalln("Error:", err)
	}
	y :=base64.StdEncoding.EncodeToString(b)
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-fino",
			Value: id.String()+"|"+name+age,
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	cookie.Value = y
	err = tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func main(){
	http.HandleFunc("/", snackWells)
	http.Handle("/favicon.ico", http.NotFoundHandler()) //needed?
	//log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}