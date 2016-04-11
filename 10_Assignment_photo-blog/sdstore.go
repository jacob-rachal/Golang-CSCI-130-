// Jacob Rachal 4-5-16 Wed.
// Fresno State CSCI 130\
//assignment 9
package mem

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
	//"log"
)

func storeDstore(m model, id string, req *http.Request) error {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Photos", id, 0, nil)

	_, err := datastore.Put(ctx, key, &m)
	if err != nil {
		(ctx, "ERROR storeDstore datastore.Put: %s", err)
		return err
	}
	return nil
}

func retrieveDstore(id string, req *http.Request) (model, error) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Photos", id, 0, nil)

	var m model
	err := datastore.Get(ctx, key, &m)
	if err != nil {
		log.Errorf(ctx, "ERROR retrieveDstore datastore.Get: %s", err)
		return m, err
	}// don't we need a case where err is nil?
	return m, nil
}

