//Jacob Rachal 1/29/2016
//CSCI 130 "Web Programing"
//Assignment 1

package main //_1_Assignment_condit_logic

import(
	"html/template"
	"log"
	"os"
	//"net/http"
)

type person struct{
	Name string
	Mark string
}
type tankOfTanks struct {
	person
	OfTheLine bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main (){
	p1 := tankOfTanks{
		person: person{
			Name: "Nike",
			Mark:  "Mark XXIII Bolo",
		},
		OfTheLine: true,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

	/*p2 := tankOfTanks{
		person: person{
			Name: "Bert The Avenger",
			Mark: "stuff",
		},
		OfTheLine: false,
	}
	err := tpl.Execute(os.Stdout, p2)
	if err != nil {
		log.Fatalln(err)
	}*/

}