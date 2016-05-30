package main//finished_questionmark
import 	(
	"crypto/hmac"
	"crypto/sha256"
)
func HmacEncode(data string) string {
	h := hmac.New(sha256.New, []byte(data+"superSecureKey"))
	return string(h.Sum(nil))
}