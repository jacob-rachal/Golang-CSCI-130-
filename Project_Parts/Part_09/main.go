// Jacob Rachal 3-24-16
// Fresno State CSCI 130
// Project Step 9
// INFO: PROJECT STEP 9 - A user should not be able to access the form to upload user data when they are not logged in.
package main//Part_09
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