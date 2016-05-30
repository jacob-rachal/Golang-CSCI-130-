package main//finished_questionmark
import "net/http"
func makeDefaultCookie() (*http.Cookie){
	result := getDefaultValue()
	return &http.Cookie{
		Name:  "session-info",
		Value: result,
		// Secure: true,
		HttpOnly: true,
	}
}