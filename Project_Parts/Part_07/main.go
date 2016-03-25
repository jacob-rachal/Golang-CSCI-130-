// Jacob Rachal 3-24-16
// Fresno State CSCI 130
// Project Step 7
// INFO: PROJECT STEP 7 - Allow the user to login.
// Store the information about whether or not a user is logged-in
// in both the "user" data type you created and in the cookie.
// Show a "logout" button when the user is logged in
package main//Part_07
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
	Username string //since we are now having users log in and out,
			// now would be a good time to change the purpose of these strings.
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