// Jacob Rachal 3-13-16
// Fresno State CSCI 130
// PROJECT PART 2 (aka STEP 2)
// INSTR; have the application write a cookie called "session-fino" with a UUID.
// The cookie should serve HttpOnly and you should have the "Secure" flag set also
// though comment the "Secure" flag out as we're not using https.
package main //Part_02
import(
	"net/http"
	"html/template"
	//"io"
	"log"
	"github.com/nu7hatch/gouuid" //was I supposed to download this repository or something? Because it is turning up red for me.
)

func handler( res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("temp.html")
	if err != nil{
		log.Fatalln("Something went wrong: ", err)
	}
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-fino",
			Value: id.String(),
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
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}