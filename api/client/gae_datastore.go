package client

import (
	"net/http"
//	"golang.org/x/net/context"
//	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"github.com/mjibson/goon"
	"log"

)

type GaeDatastoreInstance struct {
	gooninstance *goon.Goon
	Keys []*datastore.Key
}

func NewGaeDatastore(r *http.Request) *GaeDatastoreInstance {	
	I := &GaeDatastoreInstance{
		gooninstance: goon.NewGoon(r),
	}
	return I
}

func (gdi *GaeDatastoreInstance) GetOne(entitydata interface{}) error {
	return gdi.gooninstance.Get(entitydata)
}

func (gdi *GaeDatastoreInstance) GetAll(entity, filter, key string, entitydata interface{}) ([]*datastore.Key, error) {
	q := datastore.NewQuery(entity).Filter(filter, key)
	return gdi.gooninstance.GetAll(q, entitydata)
}

func (gdi *GaeDatastoreInstance) Put(entitydata interface{}) error {
	log.Print("---2--------")
	log.Print(entitydata)
	_, err := gdi.gooninstance.Put(entitydata)
	return err
}