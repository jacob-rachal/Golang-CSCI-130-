// Jacob Rachal 5/30/16 CSCI 130 Fresno State
//Part 8 of the big project (not the group one)
package main//finished_questionmark
import "net/http"

func main(){
	http.HandleFunc("/", webpage)
	http.HandleFunc("/login",loggingin)
	http.HandleFunc("/logout",loggingout)
	http.ListenAndServe(":8080",nil)
}