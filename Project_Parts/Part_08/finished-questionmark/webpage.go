package main//finished_questionmark
import(
	"net/http"
)
func webpage(res http.ResponseWriter,req *http.Request){
	obj := handleCookie(res,req)
	executeTemplate(res,obj)
}