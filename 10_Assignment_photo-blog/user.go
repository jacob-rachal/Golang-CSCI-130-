// Jacob Rachal 4-11-16
// Fresno State CSCI 130
// assignment 10 user.go
package mem

import (
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"net/http"
	"google.golang.org/appengine/log"
)

func newVisitor(req *http.Request) (*http.Cookie, error) {
	id, err := uuid.NewV4()
	if err != nil {
		ctx := appengine.NewContext(req)
		log.Errorf(ctx, "ERROR newVisitor uuid.NewV4: %s", err)// this line, the one above,
		// and the imports are the only thing to differentiate this file from the one in the previous assignment.
		return nil, err
	}
	m := initialModel(id.String())
	return makeCookie(m, req)
}

func currentVisitor(m model, id string, req *http.Request) (*http.Cookie, error) {
	return makeCookie(m, req)
}

func initialModel(id string) model {
	m := model{
		Name:  "",
		State: false,
		Pictures: []string{
			"one.jpg",
		},
		ID: id,
	}
	return m
}
