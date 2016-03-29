// Jacob Rachal 3-22-16
// Fresno State CSCI 130
// INSTR: create a web page which serves at localhost over https using TLS.
//QUESTION: do I need to get a key and certificate and include them on my github?
package main//_7_Assignment_TLS
import (
	"net/http"
	"io"
	"log"
)

func BasicServer(res http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "text/plain")
	//w.Write([]byte("This is an example server.\n"))
	io.WriteString(res, "Hello World!")
}

func main() {
	http.HandleFunc("/Hello", BasicServer)

	//go http.ListenAndServe(":8080", http.RedirectHandler("https://127.0.0.1:10443/", 301)) //was commented out

	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil) //or should it be ":443"?
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}