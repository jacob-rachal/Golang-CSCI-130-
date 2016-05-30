package main//finished_questionmark
import (
	"net/http"
	"fmt"
)

func handleCookie(res http.ResponseWriter,req *http.Request) User {
	cookie, err := req.Cookie("session-info")
	if err != nil {
		cookie = makeDefaultCookie()
		http.SetCookie(res,cookie)
	}
	if req.Method == "POST" {
		cookie.Value = updateCookie(cookie,req)
		http.SetCookie(res,cookie)
	}
	if cookieIsBad(cookie){
		fmt.Println(req.RemoteAddr,": Tried to edit their information.")
	}
	return outputUser(cookie.Value)
}