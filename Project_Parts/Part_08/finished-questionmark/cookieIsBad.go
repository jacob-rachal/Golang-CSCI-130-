package main//finished_questionmark
import (
	"net/http"
	"crypto/hmac"
	"encoding/base64"
	"strings"
)
func cookieIsBad(cookie *http.Cookie) bool {
	stringsToParse := strings.Split(cookie.Value,",")
	t, _ := base64.StdEncoding.DecodeString(stringsToParse[1])
	return !(hmac.Equal([]byte(HmacEncode(stringsToParse[0])),[]byte(t)))
}