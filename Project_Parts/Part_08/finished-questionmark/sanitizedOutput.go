package main//finished_questionmark
import(
	"encoding/base64"
)
func sanitizedOutput(input []byte) string{
	left := base64.StdEncoding.EncodeToString(input)
	delimiter := ","
	right := base64.StdEncoding.EncodeToString([]byte(HmacEncode(base64.StdEncoding.EncodeToString(input))))
	return left+delimiter+right
}