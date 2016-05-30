package main//finished_questionmark
import (
	"net/http"
	"strings"
	"fmt"
	"encoding/base64"
	"encoding/json"
)

func loggingin(res http.ResponseWriter,req *http.Request){
	cookie, err := req.Cookie("session-info")
	if err == nil {
		if cookieIsBad(cookie){
			fmt.Println(req.RemoteAddr,": Tried to edit their information at log in.")
		} else{
			var user User
			tempString := strings.Split(cookie.Value,",")[0]
			decoded, _ := base64.StdEncoding.DecodeString(tempString)
			json.Unmarshal(decoded,&user)
			user.LoggedIn = "loggedIn"
			r, _ := json.Marshal(user)
			cookie.Value = sanitizedOutput(r)
			http.SetCookie(res,cookie)

			if !cookieIsBad(cookie){
				http.Redirect(res, req, "/", http.StatusFound)
			} else {
				fmt.Println(req.RemoteAddr,": Tried to edit their information at log in.")
			}
		}
	}
}