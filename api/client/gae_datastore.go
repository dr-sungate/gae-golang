package client

import (
	"net/http"
	"fmt"
//	"golang.org/x/net/context"
//	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"github.com/mjibson/goon"
	log "gaeapp/gae-golang/api/logger"
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
	log.Info(fmt.Sprintf("%v", entitydata))
	return gdi.gooninstance.Get(entitydata)
}

func (gdi *GaeDatastoreInstance) GetAll(entity string, querymap map[string]string, limit, offset int, order string, entitydata interface{}) ([]*datastore.Key, error) {
	log.Info(fmt.Sprintf("%s / %s ", entity, querymap))
	q := datastore.NewQuery(entity)
	for key, value := range querymap {
		q = q.Filter(key, value)
	}
	if limit != 0 { 
		q = q.Limit(limit)
	}
	q = q.Offset(offset)
	if order != "" { 
		q = q.Order(order)
	}
	return gdi.gooninstance.GetAll(q, entitydata)
}

func (gdi *GaeDatastoreInstance) Put(entitydata interface{}) error {
	log.Info(fmt.Sprintf("%v", entitydata))
	_, err := gdi.gooninstance.Put(entitydata)
	return err
}

func (gdi *GaeDatastoreInstance) Delete(key *datastore.Key) error {
	log.Info(fmt.Sprintf("%v", key))
	return gdi.gooninstance.Delete(key)
}

func (gdi *GaeDatastoreInstance) GetKey(key interface{}) *datastore.Key {
	log.Info(fmt.Sprintf("%v", key))
	return gdi.gooninstance.Key(key)
}