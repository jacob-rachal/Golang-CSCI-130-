package main//finished_questionmark
import(
	"encoding/json"
	"encoding/base64"
	"strings"
)


func outputUser(Value string) User{
	var output User
	tempString := Value
	tempString = strings.Split(tempString,",")[0]
	decoded, _ := base64.StdEncoding.DecodeString(tempString)
	json.Unmarshal(decoded,&output)
	return output
}