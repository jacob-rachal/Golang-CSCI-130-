package main//finished_questionmark
import(
	"encoding/json"
	"net/http"
	"strings"
)

func updateCookie(cookie *http.Cookie,req *http.Request) string {
	tempString := cookie.Value
	tempString = strings.Split(tempString,",")[0]
	obj := outputUser(tempString)
	obj.Name= req.FormValue("name")
	obj.Age = req.FormValue("age")
	r, _ := json.Marshal(obj)
	return sanitizedOutput(r)
}