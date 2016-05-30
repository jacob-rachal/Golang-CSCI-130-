package main//finished_questionmark
import (
	"net/http"		// web server
	"html/template" // templates
)

//Runs the template.
func executeTemplate(res http.ResponseWriter,obj User){
	if obj.LoggedIn == "loggedIn" {
		t,_ := template.New("name").Parse(file1)
		//Execute the template.
		t.Execute(res, obj)
	} else{

	}
	if obj.LoggedIn == "loggedOff"{
		t,_ := template.New("name").Parse(file2)
		//Execute the template.
		t.Execute(res, obj)
	}
}
