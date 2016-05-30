package main//finished_questionmark
import(
	"encoding/json"
	"github.com/nu7hatch/gouuid"
)

func getDefaultValue() string {
	id,_ := uuid.NewV4()
	name := "NULL"
	age  := "NULL"
	loggedin := "loggedOff"
	r, _ := json.Marshal(User{id.String(),name,age,loggedin})
	return sanitizedOutput(r)
}