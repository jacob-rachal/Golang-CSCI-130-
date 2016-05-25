// Jacob Rachal 5/24/16 CSCI 130 Fresno State
// Assignment: Upload a Text File: Create a webpage that serves a form and allows the user to upload a txt file.
// You do not need to check if the file is a txt; bad programming but just trust the user to follow the instructions.
// Once a user has uploaded a txt file, copy the text from the file and display it on the webpage.
// Use req.FormFile and io.Copy to do this
package main//_5_upload_txt
import(
	"fmt"
	"net/http"
	"io"
)
func feedMe(res http.ResponseWriter, req *http.Request){
	fmt.Fprint(res,`<!DOCTYPE html>
				<html>
				<head>
					<title></title>
				</head>
					<body>
					<form method = "POST"  enctype="multipart/form-data">
						<input type="file" name="name"><br>
						<input type="submit">
					</form>
				</body>
				</html>`)
	if req.Method == "POST"{
		key := "name"
		_, hdr, err2 := req.FormFile(key)
		if err2 != nil{
			fmt.Println(err2)
		}

		rdr, err := hdr.Open()
		if err != nil{
			fmt.Println(err)
		}
		io.Copy(res,rdr)
	}
}
func main(){
	http.HandleFunc("/", feedMe)
	http.ListenAndServe(":8080", nil)
}
//Good grief, why didn't I have this done earlier? Oh right, because back then html was scary and I didn't see the distinction between ' and `.