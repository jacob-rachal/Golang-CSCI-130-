// Jacob Rachal 3-24-16
// Fresno State CSCI 130
// Project Step 8
// INFO: PROJECT STEP 8 -  Allow the user to logout.
// Show a log-in button when the user is not logged-in.
// Show a log-out button only when the user is logged in.
package main//Part_08
import(
	"net/http"
	"html/template"
	"log" //remove?
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"
)
type user struct{
	Username string
	Password string
	LogStatus bool
}
func main(){
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func getCode(data string) string { //making the HMAC
	h := hmac.New(sha256.New, []byte("H3110w0rld")) //was previously "ourkey"
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}