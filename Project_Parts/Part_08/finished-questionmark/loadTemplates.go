package main//finished_questionmark
import "io/ioutil"

var file1 string
var file2 string

func init(){
	temp, _ := ioutil.ReadFile("templates/loggedIn.htemplate")
	file1 = string(temp)
	temp, _ = ioutil.ReadFile("templates/loggedOut.htemplate")
	file2 = string(temp)
}

