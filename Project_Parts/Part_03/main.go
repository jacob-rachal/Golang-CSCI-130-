// Jacob Rachal 3-13-16
// Fresno State CSCI 130
// PROJECT PART 3 (aka STEP 3)
// INSTR: continuing to build our application, create a template which is a form.
// The form should gather the user's name and age. Store the user's name and age in the cookie.
package main //Part_03
import(
	"net/http"
	"html/template"
	//"io"
	"log"
	"github.com/nu7hatch/gouuid"
)
func snackWells( res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("temp.html")
	if err != nil{
		log.Fatalln("Something went wrong: ", err)
	}
	name:= req.FormValue("Name:")
	age:= req.FormValue("Age:")
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
	tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func main(){
	http.HandleFunc("/", snackWells)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}